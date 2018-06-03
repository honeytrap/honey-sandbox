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

	"github.com/honeytrap/honeytrap/event"

	"github.com/gorilla/websocket"
)

type web struct {
	ListenAddress string `toml:"listen"`

	messageCh chan json.Marshaler

	// Registered connections.
	connections map[*connection]bool

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func New(options ...func(*web) error) (*web, error) {
	hc := web{
		ListenAddress: "127.0.0.1:8089",

		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),

		messageCh: make(chan json.Marshaler),
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

	/*
		sh := http.FileServer(&assetfs.AssetFS{
			Asset:     assets.Asset,
			AssetDir:  assets.AssetDir,
			AssetInfo: assets.AssetInfo,
			Prefix:    assets.Prefix,
		})
	*/

	handler.HandleFunc("/ws", web.ServeWS)
	/*
		handler.Handle("/", sh)
	*/

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

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (msg Message) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}
	m["type"] = msg.Type
	m["data"] = msg.Data
	return json.Marshal(m)
}

func Data(t string, data interface{}) json.Marshaler {
	return &Message{
		Type: t,
		Data: data,
	}
}

func (web *web) Send(evt event.Event) {
	web.eventCh <- evt
}

func (web *web) ServeWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Could not upgrade connection: %s", err.Error())
		return
	}

	c := &connection{
		ws:   ws,
		web:  web,
		send: make(chan json.Marshaler, 100),
	}

	log.Info("Connection upgraded.")
	defer func() {
		c.web.unregister <- c
		c.ws.Close()

		log.Info("Connection closed")
	}()

	web.register <- c

	go c.writePump()
	c.readPump()
}
