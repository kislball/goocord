package goocord

import (
	"testing"
)

var TestingProvider = NewMapCacheProvider()

func TestMapCacheProvider_Set(t *testing.T) {
	if err := TestingProvider.Set("hello", "test", "idk"); err != nil {
		t.Error("expected MapCacheProvider.Set not to return an error")
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
		t.Error("expected TestingProvider.Get to return NotFoundError")
	}
	if v != nil {
		t.Error("expected TestingProvider.Get to return nil")
	}
}
