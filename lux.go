package logx

import (
	"golang.org/x/exp/slog"
	"sync"
)

const (
	ContextLoggerKey = "_logger_"

	ErrorStackFieldName = "stack"
)

type Rule func(record slog.Record) bool

type Target interface {
	Handle(slog.Record) error
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
//	muxHandler := ltt.New()
func New() Mux {
	return &muxHandler{
		tree: &muxEntry{
			target: &nullTarget{},
			match: func(_ slog.Record) bool {
				return true
			},
		},
		lock: sync.Mutex{},
	}
}
