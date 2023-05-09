package target

import (
	"errors"
	"github.com/nooize/ltt"
)

func Distributor(targets ...ltt.Target) (ltt.Target, error) {
	return &distributeTarget{
		targets: targets,
	}, nil
}

type distributeTarget struct {
	targets []ltt.Target
}

func (dt *distributeTarget) Handle(e ltt.Event) (out error) {
	for _, trg := range dt.targets {
		if err := trg.Handle(e); err != nil {
			// TODO fix message
			out = errors.Join(out, err)
		}
	}
	return out
}