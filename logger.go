package lwr

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

// Logger represent context for logging even
type Logger interface {

	// WithContext returns a new Context that has logger inside	it
	WithContext(context.Context) context.Context

	// With returns a new Logger with tags added to the logger.
	With(string, interface{}) Logger

	// Trace starts a new message with trace level.
	// You must call Send on the returned event in order to newEvent the event.
	Trace(string, ...interface{})

	// Debug starts a new message with debug level.
	// You must call Send on the returned event in order to newEvent the event.
	Debug(string, ...interface{})

	// Info starts a new message with info level.
	// You must call Send on the returned event in order to newEvent the event.
	Info(string, ...interface{})

	// Warn starts a new message with warn level.
	// You must callSend on the returned event in order to newEvent the event.
	Warn(string, ...interface{})

	// Error starts a new message with error level.
	// You must call Send on the returned event in order to newEvent the event.
	Error(string, ...interface{})

	// Fatal starts a new message with fatal level. The os.Exit(1) function
	// is called by the Send method, which terminates the program immediately.
	Fatal(string, ...interface{})
}

type logger struct {
	tags  *Tags
	route func(Event)
}

func (c *logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextLoggerKey, c)
}

func (c *logger) With(key string, v interface{}) Logger {
	if len(key) == 0 || v == nil {
		return c
	}
	switch reflect.TypeOf(v).Kind() {
	// functions and channels are not supported as tags
	case reflect.Func, reflect.Chan:
		return c
	// for nil values, we do not add the key to the tags
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Slice:
		if reflect.ValueOf(v).IsNil() {
			return c
		}
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.String, reflect.Struct:
	}
	return &logger{
		tags: &Tags{
			parent: c.tags,
			key:    key,
			val:    v,
		},
		route: c.route,
	}
}

func (c *logger) Trace(msg string, options ...interface{}) {
	c.send(Trace, msg, options...)
}

// Debug send a new log message with debug level.
func (c *logger) Debug(msg string, options ...interface{}) {
	c.send(Debug, msg, options...)
}

// Info send a new log message with info level.
func (c *logger) Info(msg string, options ...interface{}) {
	c.send(Info, msg, options...)
}

// Warn send a new log message with warn level.
func (c *logger) Warn(msg string, options ...interface{}) {
	c.send(Warning, msg, options...)
}

// Error send a new log message with error level.
func (c *logger) Error(msg string, options ...interface{}) {
	c.send(Error, msg, options...)
}

// Fatal send a new log with fatal level. The os.Exit(1) function
// is called by the Send method, which terminates the program immediately.
func (c *logger) Fatal(msg string, options ...interface{}) {
	c.send(Fatal, msg, options...)
}

func (c *logger) send(level Level, msg string, options ...interface{}) {
	c.route(&event{
		time:  time.Now(),
		level: level,
		tags:  c.tags,
		msg:   fmt.Sprintf(msg, options...),
	})
}
