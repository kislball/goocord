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
