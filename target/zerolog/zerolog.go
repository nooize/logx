package zerolog

import (
	"fmt"
	"github.com/nooize/ltt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewTarget(options ...Option) ltt.Target {
	t := &zerologTarget{
		logger: log.Logger,
	}
	for _, option := range options {
		option(t)
	}
	return t
}

type zerologTarget struct {
	logger zerolog.Logger
}

func (t zerologTarget) Handle(event ltt.Event) error {
	var zeroEvent *zerolog.Event
	switch event.Level() {
	case ltt.Nop, ltt.Disabled:
		return nil
	case ltt.Trace:
		zeroEvent = t.logger.Trace()
	case ltt.Debug:
		zeroEvent = t.logger.Debug()
	case ltt.Info:
		zeroEvent = t.logger.Info()
	case ltt.Warning:
		zeroEvent = t.logger.Warn()
	case ltt.Error:
		zeroEvent = t.logger.Error()
	case ltt.Fatal:
		zeroEvent = t.logger.Fatal()
	default:
		return fmt.Errorf("unknown log level: %d", event.Level())
	}
	if tags := event.Tags(); tags != nil {
		for key, val := range tags.ToMap() {
			zeroEvent = zeroEvent.Interface(key, val)
		}
	}
	zeroEvent = zeroEvent.Time(zerolog.TimestampFieldName, event.Time())
	zeroEvent.Msg(event.Message())
	return nil
}
