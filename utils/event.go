package utils

// EventEmitter represents a message bus
type EventEmitter struct {
	Handlers map[string][]interface{}
}

// Emit events an event
func (e *EventEmitter) Emit(name string, data interface{}) {
	handlers := e.Handlers[name]

	for _, v := range handlers {
		c := v.(func (data interface{}))
		go c(data)
	}
}

// AddHandler adds a new listener
func (e *EventEmitter) AddHandler(name string, listener interface{}) {
	if e.Handlers == nil {
		e.Handlers = map[string][]interface{}{}
	}

	if _, ok := e.Handlers[name]; !ok {
		e.Handlers[name] = []interface{}{}
	}
	e.Handlers[name] = append(e.Handlers[name], listener)
}

// On is an alias for AddHandler
func (e *EventEmitter) On(name string, listener interface{}) {
	e.AddHandler(name, listener)
}
