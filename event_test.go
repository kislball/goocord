package goocord

import (
	"testing"
	"time"
)

var emitter EventEmitter = EventEmitter{}

func TestEventEmitter_AddHandler(t *testing.T) {
	emitted := false
	emitter.AddHandler("test", func(test string) {
		emitted = true
	})

	emitter.Emit("test", "hello")
	time.Sleep(time.Second)
	if !emitted {
		t.Fatal("event was not emitted")
	}
}
