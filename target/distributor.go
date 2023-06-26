package target

import (
	"errors"
	"github.com/nooize/lux"
)

func Distributor(targets ...lux.Target) (lux.Target, error) {
	return &distributeTarget{
		targets: targets,
	}, nil
}

type distributeTarget struct {
	targets []lux.Target
}

func (dt *distributeTarget) Handle(e lux.Event) (out error) {
	for _, trg := range dt.targets {
		if err := trg.Handle(e); err != nil {
			// TODO fix message
			out = errors.Join(out, err)
		}
	}
	return out
}
