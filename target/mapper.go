package target

import "github.com/nooize/lwr"

func LevelMapper(next lwr.Target, mapping func(lwr.Level) lwr.Level) lwr.Target {
	trg := &levelMapperTarget{
		next:    next,
		mapping: mapping,
	}
	return trg
}

type levelMapperTarget struct {
	next    lwr.Target
	mapping func(lwr.Level) lwr.Level
}

func (ct *levelMapperTarget) Handle(e lwr.Event) error {
	// TODO implement
	return ct.next.Handle(e)
}
