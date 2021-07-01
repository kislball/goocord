package utils

import (
	"reflect"
	"sync"
)

// Utility structure used for event management
type EventEmitter struct {
	sync.RWMutex
	handlers map[string][]reflect.Value
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{RWMutex: sync.RWMutex{}}
}

// Add an event handler
func (e *EventEmitter) On(name string, handler interface{}) {
	e.Lock()
	defer e.Unlock()
	val := reflect.ValueOf(handler)

	if e.handlers == nil {
		e.handlers = map[string][]reflect.Value{}
		e.handlers[name] = []reflect.Value{}
	} else if e.handlers[name] == nil {
		e.handlers[name] = []reflect.Value{}
	}

	e.handlers[name] = append(e.handlers[name], val)
}

// Emit an event
func (e *EventEmitter) Emit(name string, data ...interface{}) {
	e.RLock()
	defer e.RUnlock()

	if e.handlers == nil {
		e.handlers = map[string][]reflect.Value{}
		e.handlers[name] = []reflect.Value{}
	} else if e.handlers[name] == nil {
		e.handlers[name] = []reflect.Value{}
	}

	values := []reflect.Value{}

	for _, v := range data {
		values = append(values, reflect.ValueOf(v))
	}

	for _, v := range e.handlers[name] {
		go v.Call(values)
	}
}
