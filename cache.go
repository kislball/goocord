package goocord

import "sync"

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
	return &MapCacheProvider{map[string]*MapCacheNamespace{}}
}

// NewMapCacheNamespace creates a new MapCacheNamespace
func NewMapCacheNamespace() *MapCacheNamespace {
	return &MapCacheNamespace{sync.RWMutex{}, map[string]interface{}{}}
}

// getNamespace gets a MapCacheNamespace from MapCacheProvider
func (c *MapCacheProvider) getNamespace(namespace string) *MapCacheNamespace {
	ns := c.namespaces[namespace]
	if ns == nil {
		c.namespaces[namespace] = NewMapCacheNamespace()
		ns = c.namespaces[namespace]
	}
	return ns
}

// Get gets a key from specific MapCacheNamespace
func (c *MapCacheProvider) Get(namespace string, key string) (interface{}, error) {
	ns := c.getNamespace(namespace)
	v := ns.data[key]
	if v == nil {
		return nil, NotFoundError
	}
	return v, nil
}

// Set sets a key in given MapCacheNamespace
func (c *MapCacheProvider) Set(namespace string, key string, value interface{}) error {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	ns.data[key] = value
	return nil
}

// Delete deletes a key pair from MapCacheNamespace
func (c *MapCacheProvider) Delete(namespace string, key string) error {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	delete(ns.data, key)
	return nil
}

// Clear clears MapCacheNamespace
func (c *MapCacheProvider) Clear(namespace string) error {
	ns := c.getNamespace(namespace)
	ns.Lock()
	defer ns.Unlock()
	ns.data = map[string]interface{}{}
	return nil
}

// GetAll gets all pairs from MapCacheNamespace
func (c *MapCacheProvider) GetAll(namespace string) (map[string]interface{}, error) {
	ns := c.getNamespace(namespace)
	return ns.data, nil
}

// Total returns total amount of pairs stored in MapCacheNamespace
func (c *MapCacheProvider) Total(namespace string) (int, error) {
	ns := c.getNamespace(namespace)
	return len(ns.data), nil
}
