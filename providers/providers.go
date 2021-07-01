package providers

import (
	"errors"
	"github.com/kislball/goocord/types/gateway"
	"github.com/kislball/goocord/utils"
)

// Providers represent a set of providers used by Client
type Providers struct {
	Cache *CacheProvider
}

var NotFoundError = errors.New("not found")
var ProviderNotReadyError = errors.New("provider is not ready")

// CacheProvider represents a cache storage.
type CacheProvider interface {
	// Get a key from Key/Value storage.
	Get(namespace string, key string) (interface{}, error)
	// Createsa new Key/Value pair in storage.
	Set(namespace string, key string, value interface{}) error
	// Delete a pair in Key/Value storage
	Delete(namespace string, key string) error
	// Clear entire namespace
	Clear(namespace string) error
	// Get all key/value pairs
	GetAll(namespace string) (map[string]interface{}, error)
	// Total amount of k/v pairs
	Total(namespace string) (int, error)
}

// RestProvider represents a requester which sends requests to Discord.
// Internally, GooCord doesn't implement ratelimiting, RestProviders are responsible for that
type RestProvider interface {
	// Set an authorization header
	UseAuth(token string)
	// Set API url
	UseAPI(url string)
	// Send a request to Discord API
	Request(method string, endpoint string, headers map[string]string, body interface{}) (*RestResponse, error)
}

// RestResponse represents a response from Discord API
type RestResponse struct {
	// HTTP status code
	StatusCode int
	// HTTP Headers
	Headers map[string]string
	// Body
	Body interface{}
}

// GatewayProvider represents a bi-directional connection between
// Discord and GooCord. A single GatewayProvider can only handle
// one shard.
type GatewayProvider interface {
	// Set the token to use
	UseToken(token string)
	// Open connection to Discord with given shard ID and total shards
	Connect(shard int, total int) error
	// Close the connection
	Close() error
	// Add an OnOpen handler
	OnOpen(func())
	// Add an OnClose handler
	OnClose(func())
	// Add an OnPacket handler
	OnPacket(func(d GatewayProviderOnPacketData))
	// Send a packet
	Send(json interface{}) error
	// Shard ID and total shards ran by this provider
	ShardInfo() [2]int
	// Sets presence
	UsePresence(presence gateway.UpdatePresence) error
	// Use intents
	UseIntents(intents utils.Flags) error
}

// Payload in OnPacket
type GatewayProviderOnPacketData interface {
	Data() interface{} // Get data
}
