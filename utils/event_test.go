package utils

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	emitter := NewEventEmitter()
	emitted := false
	val := ""

	emitter.On("hello", func (name string) {
		val = fmt.Sprintf("Hello, %s", name)
		emitted = true
	})

	emitter.Emit("hello", "Peter")

	time.Sleep(500 * time.Millisecond)

	if !emitted || val != "Hello, Peter" {
		t.Error("event was not emitted")
	}
}
