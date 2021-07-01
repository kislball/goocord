package providers

import (
	"sync"
)

// MapCacheProvider is a basic CacheProvider used by default. Uses map
// as its main storage.
type MapCacheProvider struct {
	namespaces map[string]*MapCacheNamespace
}

// MapCacheNamespace is a unit of cache used to separate different
// keys.
type MapCacheNamespace struct {
	sync.RWMutex
	data map[string]interface{}
}

// NewMapCacheProvider creates a new MapCacheProvider
func NewMapCacheProvider() *MapCacheProvider {
	return &MapCacheProvider{
		namespaces: map[string]*MapCacheNamespace{},
	}
}

// NewMapCacheNamespace creates a new MapCacheNamespace
func NewMapCacheNamespace() *MapCacheNamespace {
	return &MapCacheNamespace{
		RWMutex: sync.RWMutex{},
		data:    make(map[string]interface{}),
	}
}

// getNamespace gets a MapCacheNamespace from MapCacheProvider
func (c *MapCacheProvider) getNamespace(namespace string) (ns *MapCacheNamespace) {
	ns = c.namespaces[namespace]
	if ns == nil {
		ns = NewMapCacheNamespace()
		c.namespaces[namespace] = ns
	}
	return
}

// Get gets a key from specific MapCacheNamespace
func (c *MapCacheProvider) Get(namespace string, key string) (rv interface{}, err error) {
	ns := c.getNamespace(namespace)
	rv, ok := ns.data[key]
	if !ok {
		err = NotFoundError
	}
	return
}

// Set sets a key in given MapCacheNamespace
func (c *MapCacheProvider) Set(namespace string, key string, value interface{}) (err error) {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	ns.data[key] = value
	return
}

// Delete deletes a key pair from MapCacheNamespace
func (c *MapCacheProvider) Delete(namespace string, key string) (err error) {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	delete(ns.data, key)
	return
}

// Clear clears MapCacheNamespace
func (c *MapCacheProvider) Clear(namespace string) (err error) {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	ns.data = make(map[string]interface{})
	return
}

// GetAll gets all pairs from MapCacheNamespace
func (c *MapCacheProvider) GetAll(namespace string) (data map[string]interface{}, err error) {
	return c.getNamespace(namespace).data, nil
}

// Total returns total amount of pairs stored in MapCacheNamespace
func (c *MapCacheProvider) Total(namespace string) (res int, err error) {
	return len(c.getNamespace(namespace).data), nil
}
