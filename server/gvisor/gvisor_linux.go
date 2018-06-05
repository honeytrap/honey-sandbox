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
	"strings"
	"sync"
	"time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	gvisorlog "gvisor.googlesource.com/gvisor/pkg/log"
	"gvisor.googlesource.com/gvisor/pkg/sentry/control"
	"gvisor.googlesource.com/gvisor/pkg/sentry/kernel/auth"
	"gvisor.googlesource.com/gvisor/pkg/urpc"
	"gvisor.googlesource.com/gvisor/runsc/boot"
	"gvisor.googlesource.com/gvisor/runsc/container"
	"gvisor.googlesource.com/gvisor/runsc/specutils"

	"github.com/hpcloud/tail"

	"github.com/containernetworking/plugins/pkg/ns"

	"github.com/vishvananda/netlink"

	"golang.org/x/sys/unix"

	"github.com/honeytrap/honeybox/server/gvisor/unet"
)

type Manager interface {
	Create() (*Sandbox, error)
}

func New(options ...func(Manager) error) (Manager, error) {
	/*
		l, err := net.Listen("unix", "/tmp/echo.sock")
		if err != nil {
			fmt.Println("listen error", err.Error())
			return nil, err
		}

		go func() {
			for {
				fd, err := l.Accept()
				if err != nil {
					fmt.Println("accept error", err.Error())
					return
				}

				go func() {
					for {
						buf := make([]byte, 512)
						nr, err := fd.Read(buf)
						if err != nil {
							return
						}

						data := buf[0:nr]
						fmt.Println(string(data))
					}
				}()
			}
		}()
	*/

	d := &gvisorManager{
		gvisorConfig: gvisorConfig{
			Bridge: "lxcbr0",
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
	Root   string `toml:"root"`
}

type gvisorManager struct {
	gvisorConfig
}

type Sandbox struct {
	c *container.Container

	config *boot.Config

	mgr *gvisorManager

	Name string
	ip   string

	eventCh chan Event

	ns    ns.NetNS
	close chan struct{}
}

func (c *Sandbox) Stop() {
	c.close <- struct{}{}
}

// resolveEnvs transforms lists of environment variables into a single list of
// environment variables. If a variable is defined multiple times, the last
// value is used.
func resolveEnvs(envs ...[]string) ([]string, error) {
	// First create a map of variable names to values. This removes any
	// duplicates.
	envMap := make(map[string]string)
	for _, env := range envs {
		for _, str := range env {
			parts := strings.SplitN(str, "=", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid variable: %s", str)
			}
			envMap[parts[0]] = parts[1]
		}
	}
	// Reassemble envMap into a list of environment variables of the form
	// NAME=VALUE.
	env := make([]string, 0, len(envMap))
	for k, v := range envMap {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env, nil
}

// NewSandbox returns a new LxcSandbox from the provider.
func (d *gvisorManager) newSandbox(name string) (*Sandbox, error) {
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

	spec := DefaultSpec()

	for i := range spec.Linux.Namespaces {
		if spec.Linux.Namespaces[i].Type != specs.NetworkNamespace {
			continue
		}

		spec.Linux.Namespaces[i].Path = netNS.Path()
	}

	spec.Root.Path = d.Root

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

	platformType, _ := boot.MakePlatformType("ptrace")
	fileAccessType, _ := boot.MakeFileAccessType("proxy")
	networkType, _ := boot.MakeNetworkType("sandbox")

	rootDir := "/var/run/runsc"
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		rootDir = filepath.Join(runtimeDir, "runsc")
	}

	conf := &boot.Config{
		RootDir:        rootDir,
		Debug:          false,
		LogFilename:    "",
		LogFormat:      "json",
		DebugLogDir:    "/tmp/" + "ht_" + name,
		FileAccess:     fileAccessType,
		Overlay:        true,
		Network:        networkType,
		LogPackets:     true,
		Platform:       platformType,
		Strace:         true,
		StraceSyscalls: []string{},
		StraceLogSize:  1024,
	}

	c, err := container.Create(containerName, spec, conf, bundleDir, "" /*consoleSocket*/, pidFile)
	if err != nil {
		return nil, fmt.Errorf("error running container: %v", err)
	}

	return &Sandbox{
		c:       c,
		config:  conf,
		mgr:     d,
		ns:      netNS,
		Name:    name,
		eventCh: make(chan Event),
		ip:      cfg.IP.String(),
		close:   make(chan struct{}),
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

type Event struct {
	Text string `json:"text"`
}

type Config struct {
	Root string
}

func WithRoot(s string) func(Manager) error {
	return func(c Manager) error {
		if m, ok := (c.(*gvisorManager)); ok {
			m.Root = s
		}

		return nil
	}
}

func (d *gvisorManager) Create() (*Sandbox, error) {
	name := fmt.Sprintf("%s", nextSuffix())

	log.Debugf("Configuring new container %s", name)

	c, err := d.newSandbox(name)
	if err != nil {
		log.Errorf("Error creating container: %s", err.Error())
		return nil, fmt.Errorf("Error creating container: %s", err)
	}

	return c, nil
}

type Control struct {
	StdIn  *unet.Socket
	StdOut *unet.Socket
	StdErr *unet.Socket
}

func (s *Sandbox) Start() error {
	log.Debugf("Starting container")

	if err := s.c.Start(s.config); err != nil {
		log.Errorf("Error starting container: %s", err)
		return fmt.Errorf("Error starting container: %s", err)
	}

	filepath.Walk("/tmp/ht_"+s.Name+"/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if filepath.Ext(path) == ".gofer" {
			fmt.Println("FOUND GOFER", path)
			go func() {
				t, err := tail.TailFile(path, tail.Config{Follow: true})
				if err != nil {
					return

				}

				for line := range t.Lines {
					fmt.Println("GOFERO", line.Text)
				}
			}()
		} else if filepath.Ext(path) == ".boot" {
			// found boot!
			fmt.Println("FUOND BOOT", path)
			go func() {
				t, err := tail.TailFile(path, tail.Config{Follow: true})
				if err != nil {
					return

				}

				for line := range t.Lines {
					if !strings.Contains(line.Text, "exec") {
						continue
					}

					fmt.Println("BOOT", line.Text)
					s.eventCh <- Event{Text: line.Text}
				}
			}()
		}

		return nil
	})

	go func() {
		log.Debugf("Waiting for container to finish ")

		s.c.Wait()
	}()

	go func() {
		<-s.close

		// cleanup container
		if err := s.c.Destroy(); err != nil {
			log.Errorf("Error starting container: %s", err)
		}

		// cleanup network
		if err := UnmountNS(s.ns); err != nil {
			log.Errorf("Error unmounting namespace: %s", err)
		}
	}()

	return nil
}

func (c *Sandbox) Events() chan Event {
	return c.eventCh
}

func (c *Sandbox) Run() (*Control, error) {
	stdInReader, stdInWriter, err := unet.SocketPair(false)
	if err != nil {
		panic(err)
	}

	stdOutReader, stdOutWriter, err := unet.SocketPair(false)
	if err != nil {
		panic(err)
	}

	stdErrReader, stdErrWriter, err := unet.SocketPair(false)
	if err != nil {
		panic(err)
	}

	ctrl := Control{
		StdIn:  stdInWriter,
		StdOut: stdOutReader,
		StdErr: stdErrReader,
	}

	e := &control.ExecArgs{
		Argv:             []string{"/bin/bash"},
		WorkingDirectory: "",
		FilePayload:      urpc.FilePayload{Files: []*os.File{os.NewFile(uintptr(stdInReader.FD()), ""), os.NewFile(uintptr(stdOutWriter.FD()), ""), os.NewFile(uintptr(stdErrWriter.FD()), "")}},
		KUID:             auth.KUID(0),
		KGID:             auth.KGID(0),
		ExtraKGIDs:       []auth.KGID{},
		Capabilities:     &auth.TaskCapabilities{},
	}

	cont, err := container.Load(c.config.RootDir, "ht_"+c.Name)
	if err != nil {
		log.Fatalf("error loading sandox: %v", err)
	}

	if e.WorkingDirectory == "" {
		e.WorkingDirectory = cont.Spec.Process.Cwd
	}

	// Get the executable path, which is a bit tricky because we have to
	// inspect the environment PATH which is relative to the root path.
	// If the user is overriding environment variables, PATH may have been
	// overwritten.
	rootPath := c.c.Spec.Root.Path

	e.Envv, err = resolveEnvs(c.c.Spec.Process.Env, []string{"TERM=xterm", "LANG=en_US.UTF-8"})
	if err != nil {
		log.Fatalf("error getting environment variables: %v", err)
	}

	e.Filename, err = specutils.GetExecutablePath(e.Argv[0], rootPath, e.Envv)
	if err != nil {
		log.Fatalf("error getting executable path: %v", err)
	}

	go func() {
		if _, err := cont.Execute(e); err == io.EOF {
		} else if err != nil {
			log.Errorf("error getting processes for container: %#+v", err)
		}
	}()

	return &ctrl, nil
}

func (c *Sandbox) ensureStarted() error {
	return nil
}

// isRunning returns true/false if the container is in running state.
func (s *Sandbox) isRunning() bool {
	return s.c.Status == container.Running
}
