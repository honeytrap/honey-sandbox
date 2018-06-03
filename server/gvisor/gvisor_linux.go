/*
* Honeytrap
* Copyright (C) 2016-2017 DutchSec (https://dutchsec.com/)
*
* This program is free software; you can redistribute it and/or modify it under
* the terms of the GNU Affero General Public License version 3 as published by the
* Free Software Foundation.
*
* This program is distributed in the hope that it will be useful, but WITHOUT
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
* FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License for more
* details.
*
* You should have received a copy of the GNU Affero General Public License
* version 3 along with this program in the file "LICENSE".  If not, see
* <http://www.gnu.org/licenses/agpl-3.0.txt>.
*
* See https://honeytrap.io/ for more details. All requests should be sent to
* licensing@honeytrap.io
*
* The interactive user interfaces in modified source and object code versions
* of this program must display Appropriate Legal Notices, as required under
* Section 5 of the GNU Affero General Public License version 3.
*
* In accordance with Section 7(b) of the GNU Affero General Public License version 3,
* these Appropriate Legal Notices must retain the display of the "Powered by
* Honeytrap" logo and retain the original copyright notice. If the display of the
* logo is not reasonably feasible for technical reasons, the Appropriate Legal Notices
* must display the words "Powered by Honeytrap" and retain the original copyright notice.
 */
package gvisor

import (
	"fmt"
	"io"
	"io/ioutil"
	_ "log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	gvisorlog "gvisor.googlesource.com/gvisor/pkg/log"
	"gvisor.googlesource.com/gvisor/runsc/boot"
	"gvisor.googlesource.com/gvisor/runsc/container"
	"gvisor.googlesource.com/gvisor/runsc/specutils"

	"github.com/containernetworking/plugins/pkg/ns"

	"github.com/vishvananda/netlink"

	"golang.org/x/sys/unix"

	"github.com/honeytrap/honey-container/server/gvisor/unet"
)

type Manager interface {
	New() (*gvisorSandbox, error)
}

func New(options ...func(Manager) error) (Manager, error) {
	platformType, _ := boot.MakePlatformType("ptrace")
	fileAccessType, _ := boot.MakeFileAccessType("proxy")
	networkType, _ := boot.MakeNetworkType("sandbox")

	rootDir := "/var/run/runsc"
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		rootDir = filepath.Join(runtimeDir, "runsc")
	}

	d := &gvisorManager{
		gvisorConfig: gvisorConfig{
			Bridge: "lxcbr0",
		},
		conf: &boot.Config{
			RootDir:        rootDir,
			Debug:          false,
			LogFilename:    "",
			LogFormat:      "json",
			DebugLogDir:    "/tmp/echo.sock",
			FileAccess:     fileAccessType,
			Overlay:        true,
			Network:        networkType,
			LogPackets:     true,
			Platform:       platformType,
			Strace:         true,
			StraceSyscalls: []string{},
			StraceLogSize:  1024,
		},
	}

	for _, optionFn := range options {
		optionFn(d)
	}

	if d.Bridge == "" {
		return nil, fmt.Errorf("gVisor bridge not configured.")
	}

	gvisorlog.SetLevel(gvisorlog.Debug)

	var e gvisorlog.Emitter
	e = GoogleEmitter{}

	gvisorlog.SetTarget(e)

	return d, nil
}

type GoogleEmitter struct {
	// Emitter is the underlying emitter.
}

// Emit emits the message, google-style.
func (g GoogleEmitter) Emit(level gvisorlog.Level, timestamp time.Time, format string, args ...interface{}) {
	fmt.Printf("level=%s, time=%s, format=%s args=%#+v\n", args)
}

type gvisorConfig struct {
	Bridge string `toml:"bridge"`
}

type gvisorManager struct {
	gvisorConfig

	conf *boot.Config

	name string
	ip   string
}

type gvisorSandbox struct {
	*container.Container

	name string
	ip   string

	ns    ns.NetNS
	close chan struct{}
}

func (c *gvisorSandbox) Stop() {
	c.close <- struct{}{}
}

// NewSandbox returns a new LxcSandbox from the provider.
func (d *gvisorManager) newSandbox(name string) (*gvisorSandbox, error) {
	cfg := networkConfig{
		Name:            name,
		Interface:       "eth0",
		BridgeInterface: d.Bridge,
		IP:              net.ParseIP("10.0.3.101"),
		GatewayIP:       net.ParseIP("10.0.3.1"),
	}

	// prepare network
	netNS, err := d.setupNetwork(cfg)
	if err != nil {
		log.Errorf("Error setting up network: %s", err.Error())
		return nil, fmt.Errorf("Error setting up network: %s", err)
	}

	// prepare container
	bundleDir := getwdOrDie()

	spec, err := specutils.ReadSpec(bundleDir)
	if err != nil {
		return nil, fmt.Errorf("error reading spec: %v", err)
	}

	for i := range spec.Linux.Namespaces {
		if spec.Linux.Namespaces[i].Type != specs.NetworkNamespace {
			continue
		}

		spec.Linux.Namespaces[i].Path = netNS.Path()
	}

	containerName := fmt.Sprintf("ht_%s", name)

	consoleSocket := filepath.Join("/tmp/", containerName, "console")

	if err := os.MkdirAll(filepath.Dir(consoleSocket), 0700); err != nil {
		return nil, err
	}

	srv, err := unet.BindAndListen(consoleSocket, false)
	if err != nil {
		return nil, fmt.Errorf("error binding and listening to socket %q: %v", consoleSocket, err)
	}
	/*
		defer func() {
			srv.Close()

			os.Remove(consoleSocket)
		}()
	*/

	go func() {
		for {
			// Open the othe end of the socket.
			sock, err := srv.Accept()
			if err != nil {
				log.Fatalf("error accepting socket connection: %v", err)
				return
			}

			go func() {
				// Allow 3 fds to be received.  We only expect 1.
				r := sock.Reader(true /* blocking */)
				r.EnableFDs(1)

				b := [][]byte{{}}
				if _, err := r.ReadVec(b); err != nil && err != io.EOF {
					log.Fatalf("error reading from socket connection: %v", err)
				}

				// We should have gotten a control message.
				fds, err := r.ExtractFDs()
				if err != nil {
					log.Fatalf("error extracting fds from socket connection: %v", err)
				}

				if len(fds) != 1 {
					log.Fatalf("got %d fds from socket, wanted 1", len(fds))
				}

				// Verify that the fd is a terminal.
				if _, err := unix.IoctlGetTermios(fds[0], unix.TCGETS); err != nil {
					log.Errorf("fd is not a terminal (ioctl TGGETS got %v)", err)
				}

				fmt.Println("FINISHED console socket")
			}()
		}
	}()

	pidFile := fmt.Sprintf("/var/run/honeytrap/%s.pid", containerName)

	c, err := container.Create(containerName, spec, d.conf, bundleDir, "" /*consoleSocket*/, pidFile)
	if err != nil {
		return nil, fmt.Errorf("error running container: %v", err)
	}

	return &gvisorSandbox{
		Container: c,

		ns:    netNS,
		name:  name,
		ip:    cfg.IP.String(),
		close: make(chan struct{}),
	}, nil
}

// getwdOrDie returns the current working directory and dies if it cannot.
func getwdOrDie() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current working directory: %v", err)
	}
	return wd
}

type networkConfig struct {
	Name            string
	Interface       string
	BridgeInterface string
	IP              net.IP
	GatewayIP       net.IP
}

func (ns *networkConfig) LinkName() string {
	return fmt.Sprintf("LINK_%s", ns.Name)
}

// IsExistLinkInNS finds interface name in given namespace. if foud return true.
// otherwise false.
func IsExistLinkInNS(nsName string, linkName string) (result bool, err error) {
	var vethNs ns.NetNS
	result = false

	if nsName == "" {
		if vethNs, err = ns.GetCurrentNS(); err != nil {
			return false, err
		}
	} else {
		if vethNs, err = ns.GetNS(nsName); err != nil {
			return false, err
		}
	}

	defer vethNs.Close()
	err = vethNs.Do(func(_ ns.NetNS) error {
		link, err := netlink.LinkByName(linkName)
		if link != nil {
			result = true
		}
		result = false
		return err
	})

	return result, err
}

func (d *gvisorManager) setupNetwork(cfg networkConfig) (ns.NetNS, error) {
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name:  cfg.Interface,
			Flags: net.FlagUp,
			MTU:   1500,
		},
		PeerName: cfg.LinkName(),
	}

	if err := netlink.LinkAdd(veth); err != nil {
		return nil, err
	}

	vethNs, err := NewNS(cfg.Name)
	if err != nil {
		return nil, err
	}

	defer vethNs.Close()

	if err = netlink.LinkSetNsFd(veth, int(vethNs.Fd())); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	err = vethNs.Do(func(_ ns.NetNS) error {
		link, err := netlink.LinkByName(cfg.Interface)
		if err != nil {
			return fmt.Errorf("failed to lookup %q in %q: %v",
				cfg.LinkName(), vethNs.Path(), err)
		}

		if err := netlink.LinkSetUp(link); err != nil {
			return fmt.Errorf("failed to set %q up: %v", cfg.LinkName(), err)
		}

		mask := net.CIDRMask(24, 32)

		addr := &netlink.Addr{IPNet: &net.IPNet{IP: cfg.IP, Mask: mask}, Label: ""}
		if err := netlink.AddrAdd(link, addr); err != nil {
			return fmt.Errorf(
				"failed to add IP addr %v to %q: %v",
				addr, cfg.LinkName(), err)
		}

		// add a gateway route
		route := netlink.Route{Gw: cfg.GatewayIP, Dst: nil}
		if err := netlink.RouteAdd(&route); err != nil {
			return err
		}

		lolink, err := netlink.LinkByName("lo")
		if err != nil {
			return fmt.Errorf("failed to lookup %q in %q: %v",
				cfg.LinkName(), vethNs.Path(), err)
		}

		if err := netlink.LinkSetUp(lolink); err != nil {
			return fmt.Errorf("failed to set %q up: %v", cfg.LinkName(), err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	link, err := netlink.LinkByName(cfg.LinkName())
	if err != nil {
		return nil, err
	}

	bridge := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{
			Name: cfg.BridgeInterface,
		},
	}
	if err := netlink.LinkSetMaster(link, bridge); err != nil {
		return nil, err
	}

	if err := netlink.LinkSetUp(link); err != nil {
		return nil, fmt.Errorf("failed to set %q up: %v", cfg.LinkName(), err)
	}

	if err := os.MkdirAll(filepath.Join("/etc", "netns", filepath.Base(vethNs.Path())), 0700); err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(filepath.Join("/etc", "netns", filepath.Base(vethNs.Path()), "resolv.conf"), []byte(`nameserver 8.8.8.8`), 0700); err != nil {
		return nil, err
	}

	return vethNs, nil
}

var rand2 uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand2
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand2 = r
	randmu.Unlock()
	return fmt.Sprintf("%x", int(1e9+r%1e9))
}

func (d *gvisorManager) New() (*gvisorSandbox, error) {
	name := fmt.Sprintf("%s", nextSuffix())

	log.Debugf("Configuring new container %s", name)

	c, err := d.newSandbox(name)
	if err != nil {
		log.Errorf("Error creating container: %s", err.Error())
		return nil, fmt.Errorf("Error creating container: %s", err)
	}

	log.Debugf("Starting container")

	if err := c.Start(d.conf); err != nil {
		log.Errorf("Error starting container: %s", err)
		return nil, fmt.Errorf("Error starting container: %s", err)
	}

	go func() {
		log.Debugf("Waiting for container to finish ")

		c.Wait()
	}()

	go func() {
		<-c.close

		// cleanup container
		if err := c.Destroy(); err != nil {
			log.Errorf("Error starting container: %s", err)
		}

		// cleanup network
		if err := UnmountNS(c.ns); err != nil {
			log.Errorf("Error unmounting namespace: %s", err)
		}
	}()

	return c, nil
}

func (c *gvisorSandbox) ensureStarted() error {
	/*
		if c.isFrozen() {
			return c.unfreeze()
		}

		if c.isStopped() {
			return c.start()
		}

		// settle will fill the container with ip address and interface
		return c.settle()
	*/
	return nil
}

// isRunning returns true/false if the container is in running state.
func (c *gvisorSandbox) isRunning() bool {
	return c.Status == container.Running
}
