package logx

import (
	"context"
	"golang.org/x/exp/slog"
	"sync"
)

// Mux is a log event multiplexer that route log events to one or multiple log targets
// based on predefined rules,
type Mux interface {
	slog.Handler
	AppendHandler(slog.Handler, Rule) Mux
	Logger() *slog.Logger
}

type muxHandler struct {
	tree *muxEntry
	lock sync.Mutex
}

// implement slog.Handler interface

func (m *muxHandler) Enabled(ctx context.Context, lev slog.Level) bool {
	return true
}

func (m *muxHandler) Handle(ctx context.Context, rec slog.Record) error {
	entry := m.tree
	for entry != nil {
		if entry.match(rec) {
			return entry.target.Handle(ctx, rec)
		}
		entry = entry.next
	}
	return nil
}

func (m *muxHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return m
}

func (m *muxHandler) WithGroup(name string) slog.Handler {
	return m
}

// implement Mux interface

func (m *muxHandler) Logger() *slog.Logger {
	return slog.New(m)
}

func (m *muxHandler) AppendHandler(hand slog.Handler, rule Rule) Mux {
	if hand != nil && rule != nil {
		m.lock.Lock()
		m.tree = &muxEntry{
			next:   m.tree,
			target: hand,
			match:  rule,
		}
		m.lock.Unlock()
	}
	return m
}

type muxEntry struct {
	next   *muxEntry
	target slog.Handler
	match  Rule
}
