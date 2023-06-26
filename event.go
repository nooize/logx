package lux

import "time"

type Event interface {
	Time() time.Time
	Level() Level
	Message() string
	Tags() *Tags
}

type EventMutator func(Event) (Level, string, *Tags)

type event struct {
	time  time.Time
	level Level
	tags  *Tags
	msg   string
}

func (e *event) Time() time.Time {
	return e.time
}

func (e *event) Level() Level {
	return e.level
}

func (e *event) Message() string {
	return e.msg
}

func (e *event) Tags() *Tags {
	return e.tags
}
