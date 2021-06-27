package goocord

import "github.com/gorilla/websocket"

// WebSocketGatewayProvider is a basic GatewayProvider used by default.
// Uses WS to communicate with Discord's gateway
type WebSocketGatewayProvider struct {
	dialer websocket.Dialer
	Conn   websocket.Conn
	Token  string
	EventEmitter
}

// UseToken sets a token to use
func (w *WebSocketGatewayProvider) UseToken(token string) {
	w.Token = token
}

func (w *WebSocketGatewayProvider) Connect(shard int, total int) {}

func (w *WebSocketGatewayProvider) OnOpen(handler func ()) {
	w.AddHandler("open", handler)
}

func (w *WebSocketGatewayProvider) OnClose(handler func()) {
	w.AddHandler("close", handler)
}

func (w *WebSocketGatewayProvider) OnPacket(handler func(message interface{})) {
	w.AddHandler("packet", handler)
}

func (w *WebSocketGatewayProvider) Close() {
	w.Conn.Close()
	w.Emit("close")
}
