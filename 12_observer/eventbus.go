package observer

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	mu sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		mu: sync.Mutex{},
	}
}

func (b *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := b.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	b.handlers[topic] = handler

	return nil
}

func (b *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := b.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}