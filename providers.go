package goocord

import (
	"errors"
)

// Providers represent a set of providers used by Client
type Providers struct {
	Cache *CacheProvider
}

var NotFoundError = errors.New("not found")

// CacheProvider represents a cache storage.
// Using interface{} instead of structures(e.g. Guild or Member),
// because remote cache storages(e.g. redis or memcached) cannot properly
// handle Go structures.
type CacheProvider interface {
	Get(namespace string, key string) (interface{}, error)     // Get gets a key from Key/Value storage.
	Set(namespace string, key string, value interface{}) error // Set creates a new Key/Value pair in storage.
	Delete(namespace string, key string) error                 // Delete deletes a pair in Key/Value storage
	Clear(namespace string) error                              // Clear clears entire namespace
	GetAll(namespace string) (map[string]interface{}, error)   // GetAll gets all key/value pairs
	Total(namespace string) (int, error)                       // Total returns total amount of k/v pairs
}

// RestProvider represents a requester which sends requests to Discord.
// Internally, GooCord doesn't implement ratelimiting, RestProviders are responsible for that
type RestProvider interface {
	UseToken(token string)                                                                                     // UseToken sets a token to use
	UseAPI(url string)                                                                                         // UseAPI sets an API url
	Request(method string, endpoint string, headers map[string]string, body interface{}) (RestResponse, error) // Request sends a request to Discord API
}

// RestResponse represents a response from Discord API
type RestResponse struct {
	StatusCode int               // HTTP status code
	Headers    map[string]string // HTTP Headers
	Body       interface{}       // Body
}
