package engine

import "slices"

// Signal[T] represents an event that nodes can emit and other nodes can listen to
// T is the type of the value that will be emitted with the signal
type Signal[T any] struct {
	handlers []func(T)
}

// NewSignal creates a new signal instance
func NewSignal[T any]() *Signal[T] {
	return &Signal[T]{
		handlers: make([]func(T), 0),
	}
}

// Connect adds a handler function to be called when the signal is emitted
func (s *Signal[T]) Connect(handler func(T)) {
	s.handlers = append(s.handlers, handler)
}

// Emit triggers the signal with the given value
func (s *Signal[T]) Emit(value T) {
	for _, handler := range s.handlers {
		handler(value)
	}
}

// Disconnect removes a handler from the signal
func (s *Signal[T]) Disconnect(handler func(T)) {
	for i, h := range s.handlers {
		if &h == &handler {
			s.handlers = slices.Delete(s.handlers, i, i+1)
			return
		}
	}
}
