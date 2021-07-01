package providers

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kislball/goocord/types"
	"github.com/kislball/goocord/types/gateway"
	"github.com/kislball/goocord/utils"
	"io/ioutil"
	"runtime"
)

// WebSocketGatewayProvider is a basic GatewayProvider used by default.
// Uses WS to communicate with Discord's gateway
type WebSocketGatewayProvider struct {
	dialer          *websocket.Dialer       // utility
	Conn            *websocket.Conn         // active connection
	Token           string                  // token used
	Shard           int                     // shard id
	Shards          int                     // total shards passed in IDENTIFY
	Ready           bool                    // whether the provider is ready
	Presence        *gateway.UpdatePresence // Current client's presence
	Intents         utils.Flags             // Intents used
	Sequence        *int                     // Sequence
	SessionID       *int                     // Session's id used for resume
	Zlib bool // Use zlib compression or not
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
	conn, _, err := w.dialer.Dial(types.EndpointGateway(true, 9, "json"), nil)
	w.Conn = conn
	go w.readLoop()
	w.Ready = true
	return
}

func (w *WebSocketGatewayProvider) readLoop() {
	for {
		//t, msg, _ := w.Conn.NextReader()

		//z, _ := ioutil.ReadAll(msg)
		//m, _ := w.Decode(z, t)
		//fmt.Println(t == websocket.BinaryMessage)
	}
}

// Close aborts the connection
func (w *WebSocketGatewayProvider) Close() error {
	w.Conn.Close()
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

func (w *WebSocketGatewayProvider) UseZlib() {
	if w.Ready {
		panic("tried to set zlib while bot is running")
	}
	w.Zlib = true

}

// Get heartbeat payload
func (w *WebSocketGatewayProvider) GetHeartbeat() gateway.HeartbeatPayload {
	return gateway.HeartbeatPayload{
		Payload: gateway.Payload{
			Opcode: 1,
		},
		Data: *w.Sequence,
	}
}

// Get identify payload
func (w *WebSocketGatewayProvider) GetIdentify() gateway.Identify {
	shards := w.ShardInfo()
	abc := false

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
		Compress: &abc,
	}
}

// Send heartbeat to gateway
func (w *WebSocketGatewayProvider) Beat() error {
	return w.Send(w.GetHeartbeat())
}

// Send identify payload
func (w *WebSocketGatewayProvider) Identify() error {
	return w.Send(w.GetIdentify())
}

// Decode some message
func (w *WebSocketGatewayProvider) Decode(data []byte, messageType int) (res interface{}, err error) {
	fmt.Println("hi")
	if messageType == websocket.TextMessage {
		fmt.Println("text")
		res, err = utils.Parse(string(data))
		return
	} else if messageType == websocket.BinaryMessage {
		fmt.Println("binary")
		z, err := zlib.NewReader(bytes.NewBuffer(data))
		dec, err := ioutil.ReadAll(z)
		res, err = utils.Parse(string(dec))
		fmt.Println("hey..")

		defer func() {
			z.Close()
		}()

		return res, err
	} else {
		return nil, errors.New("unsupported message type")
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
