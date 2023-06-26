package lux

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// Nop disables the logger.
	Nop Level = iota
	// Debug defines debug log level.
	Debug
	// Info defines info log level.
	Info
	// Warning defines warn log level.
	Warning
	// Error defines error log level.
	Error
	// Fatal defines fatal log level.
	Fatal

	// Trace defines trace log level.
	Trace Level = -1
	// Values less than TraceLevel are handled as numbers.

	// LevelNop is the value used for disabled logger.
	LevelNop = "-"
	// LevelDebug is the value used for the debug level field.
	LevelDebug = "debug"
	// LevelInfo is the value used for the info level field.
	LevelInfo = "info"
	// LevelWarning is the value used for the warn level field.
	LevelWarning = "warning"
	// LevelError is the value used for the error level field.
	LevelError = "error"
	// LevelFatal is the value used for the fatal level field.
	LevelFatal = "fatal"
)

// Level defines log levels
type Level int8

func (l Level) String() string {
	switch l {
	case Debug:
		return LevelDebug
	case Info:
		return LevelInfo
	case Warning:
		return LevelWarning
	case Error:
		return LevelError
	case Fatal:
		return LevelFatal
	case Nop:
		return LevelNop
	}
	return strconv.Itoa(int(l))
}

// UnmarshalText implements encoding.TextUnmarshaler to allow for easy reading from toml/yaml/json formats
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errors.New("can't unmarshal a nil *Level")
	}
	var err error
	*l, err = ParseLevel(string(text))
	return err
}

// MarshalText implements encoding.TextMarshaler to allow for easy writing into toml/yaml/json formats
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) (Level, error) {
	switch {
	case strings.EqualFold(levelStr, LevelDebug):
		return Debug, nil
	case strings.EqualFold(levelStr, LevelInfo):
		return Info, nil
	case strings.EqualFold(levelStr, LevelWarning):
		return Warning, nil
	case strings.EqualFold(levelStr, LevelError):
		return Error, nil
	case strings.EqualFold(levelStr, LevelFatal):
		return Fatal, nil
	case strings.EqualFold(levelStr, LevelNop):
		return Nop, nil
	}
	i, err := strconv.Atoi(levelStr)
	if err != nil {
		return Nop, fmt.Errorf("unknown Level: '%s', set to Nop", levelStr)
	}
	if i > 127 || i < -128 {
		return Nop, fmt.Errorf("Out-Of-Bounds Level: '%d', set to Nop", i)
	}
	return Level(i), nil
}
