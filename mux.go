package lwr

import (
	"context"
	"sync"
)

// Mux is a log event multiplexer that route log events to one or multiple log targets
// based on predefined rules,
type Mux interface {
	Ctx(ctx context.Context) Logger
	Append(Target, Rule) Mux
	Log() Logger
}

type mux struct {
	tree *muxEntry
	lock sync.Mutex
}

func (m *mux) Ctx(ctx context.Context) Logger {
	if v := ctx.Value(ContextLoggerKey); v != nil {
		if l, ok := v.(Logger); ok {
			return l
		}
	}
	return m.Log()
}

func (m *mux) Append(target Target, rule Rule) Mux {
	if target != nil && rule != nil {
		m.lock.Lock()
		m.tree = &muxEntry{
			next:   m.tree,
			target: target,
			match:  rule,
		}
		m.lock.Unlock()
	}
	return m
}

func (m *mux) Log() Logger {
	return &logger{
		route: m.route,
	}
}

func (m *mux) route(e Event) {
	entry := m.tree
	for entry != nil {
		if entry.match(e) {
			go entry.target.Handle(e)
			return
		}
		entry = entry.next
	}
}

type muxEntry struct {
	next   *muxEntry
	target Target
	match  Rule
}
