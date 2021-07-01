package providers

import (
	"testing"
)

var TestingProvider = NewMapCacheProvider()

func TestMapCacheProvider_Set(t *testing.T) {
	if err := TestingProvider.Set("hello", "test", "idk"); err != nil {
		t.Error("expected MapCacheProvider.Set not to return an error")
	}

	v, err := TestingProvider.Get("hello", "test")
	if v != "idk" || err != nil {
		t.Error("expected MapCacheProvider.Set to create a new k/v pair")
	}
}

func TestMapCacheProvider_Get(t *testing.T) {
	v, err := TestingProvider.Get("hello", "test")
	if err != nil {
		t.Error("expected MapCacheProvider.Get not to return an error")
	}
	if v != "idk" {
		t.Error("expected MapCacheProvider.Get to return \"idk\"")
	}

	v, err = TestingProvider.Get("allo", "no")
	if err != NotFoundError {
		t.Error("expected MapCacheProvider.Get to return NotFoundError")
	}
	if v != nil {
		t.Error("expected MapCacheProvider.Get to return nil")
	}
}

func TestMapCacheProvider_Total(t *testing.T) {
	if i, err := TestingProvider.Total("hello"); i != 1 || err != nil {
		t.Error("expected MapCacheProvider.Total to return (1, nil)")
	}
}

func TestMapCacheProvider_GetAll(t *testing.T) {
	m, err := TestingProvider.GetAll("hello")
	if err != nil {
		t.Error("expected MapCacheProvider.GetAll not to return an error")
	}

	if m["test"] != "idk" {
		t.Error("expected test in hello to be idk")
	}
}

func TestMapCacheProvider_Delete(t *testing.T) {
	TestingProvider.Set("hello", "test", "idk")
	if err := TestingProvider.Delete("hello", "test"); err != nil {
		t.Error("expected MapCacheProvider not to return an error")
	}

	v, err := TestingProvider.Get("hello", "test")
	if v != nil || err != NotFoundError {
		t.Error("expected MapCacheProvider.Delete to delete a k/v pair")
	}
}

func TestMapCacheProvider_Clear(t *testing.T) {
	TestingProvider.Set("hello", "test", "idk")
	if err := TestingProvider.Delete("hello", "test"); err != nil {
		t.Error("expected MapCacheProvider.Clear not to return an error")
	}

	if i, _ := TestingProvider.Total("hello"); i != 0 {
		t.Error("expected MapCacheProvider.Clear to clear collection")
	}
}
