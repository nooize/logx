package log

import (
	"context"
	"github.com/nooize/lwr"
	"github.com/nooize/lwr/rule"
	"github.com/nooize/lwr/target"
)

const (
	ContextHubKey    = "_ltt_hub_"
	ContextLoggerKey = "_ltt_logger_"

	ErrorStackFieldName = "stack"
)

var (
	rootMux lwr.Mux
)

func Append(target lwr.Target, rule lwr.Rule) {
	rootMux.Append(target, rule)
}

func Log(ctx context.Context) lwr.Logger {
	return rootMux.Ctx(ctx)
}

func With(key string, v interface{}) lwr.Logger {
	return rootMux.Log().With(key, v)
}

// Trace send a new log message with trace level.
func Trace(msg string, options ...interface{}) {
	rootMux.Log().Trace(msg, options...)
}

// Debug send a new log message with debug level.
func Debug(msg string, options ...interface{}) {
	rootMux.Log().Debug(msg, options...)
}

// Info send a new log message with info level.
func Info(msg string, options ...interface{}) {
	rootMux.Log().Info(msg, options...)
}

// Warn send a new log message with warn level.
func Warn(msg string, options ...interface{}) {
	rootMux.Log().Warn(msg, options...)
}

// Error send a new log message with error level.
func Error(msg string, options ...interface{}) {
	rootMux.Log().Error(msg, options...)
}

// Fatal send a new log with fatal level. The os.Exit(1) function
// is called by the Send method, which terminates the program immediately.
func Fatal(msg string, options ...interface{}) {
	rootMux.Log().Fatal(msg, options...)
}

func init() {
	rootMux = lwr.New()
	rootMux.Append(target.ToStderr(), rule.True())
}
