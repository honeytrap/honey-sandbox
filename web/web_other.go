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
package web

import (
	"encoding/json"
	"net/http"
	"time"

	logging "github.com/op/go-logging"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/websocket"
	assets "github.com/honeytrap/honeybox-web"
	"github.com/honeytrap/honeybox/server/gvisor"
)

var log = logging.MustGetLogger("honeyboxr/web")

type web struct {
	ListenAddress string `toml:"listen"`

	messageCh chan *message

	// Registered connections.
	connections map[*connection]bool

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func New(options ...func(*web) error) (*web, error) {
	hc := web{
		ListenAddress: "0.0.0.0:8091",

		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),

		messageCh: make(chan *message),
	}

	for _, optionFn := range options {
		if err := optionFn(&hc); err != nil {
			return nil, err
		}
	}

	return &hc, nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (web *web) Start() {
	handler := http.NewServeMux()

	server := &http.Server{
		Addr:    web.ListenAddress,
		Handler: handler,
	}

	handler.HandleFunc("/ws", web.ServeWS)

	handler.Handle("/", http.FileServer(&assetfs.AssetFS{
		Asset:     assets.Asset,
		AssetDir:  assets.AssetDir,
		AssetInfo: assets.AssetInfo,
		Prefix:    assets.Prefix,
	}))

	go web.run()

	go func() {
		log.Infof("Web interface started: %s", web.ListenAddress)

		server.ListenAndServe()
	}()
}

func (web *web) run() {
	for {
		select {
		case c := <-web.register:
			web.connections[c] = true
		case c := <-web.unregister:
			if _, ok := web.connections[c]; ok {
				delete(web.connections, c)

				close(c.send)
			}
		case msg := <-web.messageCh:
			for c := range web.connections {
				c.send <- msg
			}
		}
	}
}

type message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (msg *message) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	m["type"] = msg.Type
	m["payload"] = msg.Payload

	return json.Marshal(m)
}

func Data(data interface{}) interface{} {
	m := map[string]interface{}{}
	m["data"] = data
	return m
}

func Error(description string) *message {
	return &message{
		Type: "error",
		Payload: struct {
			Description string
		}{
			Description: description,
		},
	}
}

func Message(t string, payload interface{}) *message {
	return &message{
		Type:    t,
		Payload: payload,
	}
}

func (web *web) ServeWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Could not upgrade connection: %s", err.Error())
		return
	}

	log.Info("Connection upgraded.")

	c := &connection{
		ws:  ws,
		web: web,

		send: make(chan *message, 100),
	}

	web.register <- c

	// start new container
	mngr, err := gvisor.New(
		gvisor.WithRoot("/chroot"),
	)
	if err != nil {
		c.send <- Error(err.Error())
		return
	}

	sndbox, err := mngr.Create()
	if err != nil {
		c.send <- Error(err.Error())
		return
	}

	sndbox.Start()

	defer sndbox.Stop()

	c.send <- Message("info", struct {
		Name string `json:"name"`
	}{
		Name: sndbox.Name,
	})

	defer func() {
		c.web.unregister <- c
		c.ws.Close()

		log.Info("Connection closed")
	}()

	// watch logs
	go func() {
		for event := range sndbox.Events() {
			c.send <- Message("event", event)
			c.send <- Error("BLA")
		}
	}()

	// start new process
	ctrl, err := sndbox.Run()
	if err != nil {
		c.send <- Error(err.Error())
		return
	}

	c.ctrl = ctrl

	go func() {
		buff := make([]byte, 1024)
		for {
			n, err := c.ctrl.StdOut.Read(buff)
			if err != nil {
				log.Error(err.Error())
				break
			}

			ws.SetWriteDeadline(time.Now().Add(writeWait))

			c.send <- Message("stdout", Data(buff[:n]))
		}
	}()

	go func() {
		buff := make([]byte, 1024)
		for {
			n, err := c.ctrl.StdErr.Read(buff)
			if err != nil {
				log.Error(err.Error())
				break
			}

			ws.SetWriteDeadline(time.Now().Add(writeWait))

			c.send <- Message("stderr", Data(buff[:n]))
		}
	}()

	go c.writePump()
	c.readPump()
}

func (web *web) ServeConsoleWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Could not upgrade connection: %s", err.Error())
		return
	}

	log.Info("Connection upgraded.")

	defer func() {
		ws.Close()

		log.Info("Connection closed")
	}()

	ws.SetReadLimit(maxMessageSize)
}
