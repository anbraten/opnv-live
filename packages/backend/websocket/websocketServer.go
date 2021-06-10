package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kiel-live/kiel-live/backend/hub"
)

type WebsocketServer struct {
	// Invoked upon authentication, can be used to enforce access control.
	CanConnect func(authData map[string]interface{}) bool

	// Invoked upon channel subscription, can be used to enforce access control
	// for channels.
	CanSubscribe func(authData map[string]interface{}, channel string) bool

	// Invoked upon channel publish, can be used to enforce access control
	// for channels.
	CanPublish func(authData map[string]interface{}, channel string) bool

	// Can be set to allow CORS requests.
	CheckOrigin func(r *http.Request) bool

	// Can be used to configure buffer sizes etc.
	// See http://godoc.org/github.com/gorilla/websocket#Upgrader
	Upgrader websocket.Upgrader

	hub *hub.Hub
}

func NewWebsocketServer(hub *hub.Hub) *WebsocketServer {
	return &WebsocketServer{
		hub: hub,
	}
}

func (server *WebsocketServer) WebsocketEndpoint(w http.ResponseWriter, r *http.Request) {
	newWebsocketConnection(w, r, server)
}