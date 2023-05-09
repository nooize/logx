package target

import "github.com/nooize/ltt"

func LevelMapper(next ltt.Target, mapping func(ltt.Level) ltt.Level) ltt.Target {
	trg := &levelMapperTarget{
		next:    next,
		mapping: mapping,
	}
	return trg
}

type levelMapperTarget struct {
	next    ltt.Target
	mapping func(ltt.Level) ltt.Level
}

func (ct *levelMapperTarget) Handle(e ltt.Event) error {
	// TODO implement
	return ct.next.Handle(e)
}
