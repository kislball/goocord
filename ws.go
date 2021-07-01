package goocord

import (
	"github.com/gorilla/websocket"
	"github.com/kislball/goocord/types"
	"github.com/kislball/goocord/types/gateway"
	"github.com/kislball/goocord/utils"
	"net/http"
	"runtime"
)

// WebSocketGatewayProvider is a basic GatewayProvider used by default.
// Uses WS to communicate with Discord's gateway
type WebSocketGatewayProvider struct {
	utils.EventEmitter
	dialer    *websocket.Dialer       // utility
	Conn      *websocket.Conn         // active connection
	Token     string                  // token used
	Shard     int                     // shard id
	Shards    int                     // total shards passed in IDENTIFY
	Ready     bool                    // whether the provider is ready
	Presence  *gateway.UpdatePresence // Current client's presence
	Intents   utils.Flags             // Intents used
	Sequence  int                     // Sequence
	SessionID int                     // Session's id used for resume
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
	conn, _, err := w.dialer.Dial(types.EndpointGateway, http.Header{})
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
func (w *WebSocketGatewayProvider) OnPacket(handler func(message GatewayProviderOnPacketData)) {
	w.AddHandler("packet", handler)
}

// Close aborts the connection
func (w *WebSocketGatewayProvider) Close() error {
	w.Conn.Close()
	w.Emit("close", nil)
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

// Set presence
func (w *WebSocketGatewayProvider) UsePresence(presence gateway.UpdatePresence) (err error) {
	w.Presence = &presence
	if w.Ready {
		err = w.Send(presence)
	}
	return
}

// Set intents
func (w *WebSocketGatewayProvider) UseIntents(intents utils.Flags) error {
	if w.Ready {
		panic("tried to set intents while bot is running")
	}
	w.Intents = intents
	return nil
}

// Get heartbeat payload
func (w *WebSocketGatewayProvider) GetHeartbeat() gateway.HeartbeatPayload {
	return gateway.HeartbeatPayload{
		Payload: gateway.Payload{
			Opcode: 1,
		},
		Data: w.Sequence,
	}
}

// Get identify payload
func (w *WebSocketGatewayProvider) GetIdentify() gateway.Identify {
	shards := w.ShardInfo()

	return gateway.Identify{
		Token:   w.Token,
		Intents: w.Intents.Flags,
		Properties: gateway.IdentifyProperties{
			OS:      "Goocord",
			Browser: runtime.GOOS,
			Device:  "Goocord",
		},
		Presence: w.Presence,
		Shard:    &shards,
	}
}

// Data given to the OnPacket hook
type WebSocketGatewayProviderOnPacketData struct {
	data interface{}
}

// Get data
func (w *WebSocketGatewayProviderOnPacketData) Data() interface{} {
	return w.data
}
