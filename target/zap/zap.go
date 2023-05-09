package zap

import (
	"fmt"
	"github.com/nooize/ltt"
	"go.uber.org/zap"
)

type ZapLogger interface {
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Fatal(string, ...zap.Field)
}

func NewTarget(logger ZapLogger) (ltt.Target, error) {
	return &zapTarget{
		logger: logger,
	}, nil
}

type zapTarget struct {
	logger ZapLogger
}

func (t zapTarget) Handle(event ltt.Event) error {
	fields := make([]zap.Field, 0)
	if tags := event.Tags(); tags != nil {
		for key, val := range tags.ToMap() {
			fields = append(fields, zap.Any(key, val))
		}
	}
	switch event.Level() {
	case ltt.Nop, ltt.Disabled, ltt.Trace:
		return nil
	case ltt.Debug:
		t.logger.Debug(event.Message(), fields...)
	case ltt.Info:
		t.logger.Info(event.Message(), fields...)
	case ltt.Warning:
		t.logger.Warn(event.Message(), fields...)
	case ltt.Error:
		t.logger.Error(event.Message(), fields...)
	case ltt.Fatal:
		t.logger.Fatal(event.Message(), fields...)
	default:
		return fmt.Errorf("unknown log level: %d", event.Level())
	}
	return nil
}
