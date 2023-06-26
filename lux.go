package lwr

import (
	"sync"
)

const (
	ContextLoggerKey = "_logger_"

	ErrorStackFieldName = "stack"
)

type Rule func(Event) bool

type Target interface {
	Handle(Event) error
}

// New create new log multiplexer that route log events to one or multiple log targets
//
// Mux always have a default target with conditions:
//   - target: drop all log messages
//   - rule: match all log messages
//
// this route always reached when no other route matched
// Example:
//
//	mux := ltt.New()
func New() Mux {
	return &mux{
		tree: &muxEntry{
			target: &nullTarget{},
			match: func(_ Event) bool {
				return true
			},
		},
		lock: sync.Mutex{},
	}
}
