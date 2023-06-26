package target

import "github.com/nooize/lux"

func LevelMapper(next lux.Target, mapping func(lux.Level) lux.Level) lux.Target {
	trg := &levelMapperTarget{
		next:    next,
		mapping: mapping,
	}
	return trg
}

type levelMapperTarget struct {
	next    lux.Target
	mapping func(lux.Level) lux.Level
}

func (ct *levelMapperTarget) Handle(e lux.Event) error {
	// TODO implement
	return ct.next.Handle(e)
}
