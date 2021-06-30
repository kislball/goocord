package goocord

import (
	"github.com/kislball/goocord/types/gateway"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketGatewayProvider is a basic GatewayProvider used by default.
// Uses WS to communicate with Discord's gateway
type WebSocketGatewayProvider struct {
	dialer *websocket.Dialer // utility
	Conn   *websocket.Conn   // active connection
	Token  string            // token used
	EventEmitter
	Shard    int                   // shard id
	Shards   int                   // total shards passed in IDENTIFY
	Ready    bool                  // whether the provider is ready
	Presence gateway.UpdatePresence // Current client's presence
}

// UseToken sets a token to use
func (w *WebSocketGatewayProvider) UseToken(token string) {
	if w.Ready {
		panic("tried to set token while bot is running")
	}
	w.Token = token
}

// Connect instantiates connection to Discord
func (w *WebSocketGatewayProvider) Connect(shard int, total int) (err error) {
	w.Shard = shard
	w.Shards = total

	w.dialer = websocket.DefaultDialer
	conn, _, err := w.dialer.Dial(EndpointGateway, http.Header{})
	w.Conn = conn
	return
}

// OnOpen adds open event handler
func (w *WebSocketGatewayProvider) OnOpen(handler func()) {
	w.AddHandler("open", handler)
}

// OnClose adds close event handler
func (w *WebSocketGatewayProvider) OnClose(handler func()) {
	w.AddHandler("close", handler)
}

// OnPacket adds packet event handler
func (w *WebSocketGatewayProvider) OnPacket(handler func(message interface{})) {
	w.AddHandler("packet", handler)
}

// Close aborts the connection
func (w *WebSocketGatewayProvider) Close() error {
	w.Conn.Close()
	w.Emit("close")
	return nil
}

// Send sends data to websocket
func (w *WebSocketGatewayProvider) Send(json interface{}) error {
	if !w.Ready {
		return ProviderNotReadyError
	}

	return w.Conn.WriteJSON(json)
}

// ShardInfo returns information about shards running
func (w *WebSocketGatewayProvider) ShardInfo() [2]int {
	return [2]int{w.Shard, w.Shards}
}

func (w *WebSocketGatewayProvider) UsePresence(presence gateway.UpdatePresence) (err error) {
	w.Presence = presence
	if w.Ready {
		err = w.Send(presence)
	}
	return
}
