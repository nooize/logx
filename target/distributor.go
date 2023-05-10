package target

import (
	"errors"
	"github.com/nooize/lwr"
)

func Distributor(targets ...lwr.Target) (lwr.Target, error) {
	return &distributeTarget{
		targets: targets,
	}, nil
}

type distributeTarget struct {
	targets []lwr.Target
}

func (dt *distributeTarget) Handle(e lwr.Event) (out error) {
	for _, trg := range dt.targets {
		if err := trg.Handle(e); err != nil {
			// TODO fix message
			out = errors.Join(out, err)
		}
	}
	return out
}
