package goocord

import "reflect"

// EventEmitter represents a message bus
type EventEmitter struct {
	Handlers map[string][]reflect.Value
}

// Emit events an event
func (e *EventEmitter) Emit(name string, data ...interface{}) {
	handlers := e.Handlers[name]

	var args []reflect.Value

	for _, d := range data {
		args = append(args, reflect.ValueOf(d))
	}

	for _, v := range handlers {
		go v.Call(args)
	}
}

// AddHandler adds a new listener
func (e *EventEmitter) AddHandler(name string, listener interface{}) {
	if e.Handlers == nil {
		e.Handlers = map[string][]reflect.Value{}
	}

	if _, ok := e.Handlers[name]; !ok {
		e.Handlers[name] = []reflect.Value{}
	}
	e.Handlers[name] = append(e.Handlers[name], reflect.ValueOf(listener))
}

// On is an alias for AddHandler
func (e *EventEmitter) On(name string, listener interface{}) {
	e.AddHandler(name, listener)
}
