package target

import (
	"github.com/nooize/lux"
	"sync"
)

func Rotator(targets ...lux.Target) lux.Target {
	trg := &rotateTarget{
		targets: targets,
		cursor:  0,
		lock:    sync.Mutex{},
	}
	if len(targets) == 0 {
		trg.cursor = -1
	}
	return trg
}

type rotateTarget struct {
	targets []lux.Target
	cursor  int
	lock    sync.Mutex
}

func (rt *rotateTarget) Handle(e lux.Event) error {
	if rt.cursor < 0 {
		return nil
	}
	rt.lock.Lock()
	target := rt.targets[rt.cursor]
	rt.cursor++
	if rt.cursor >= len(rt.targets) {
		rt.cursor = 0
	}
	rt.lock.Unlock()
	return target.Handle(e)
}
