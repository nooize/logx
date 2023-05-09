package target

import (
	"github.com/nooize/ltt"
	"sync"
)

func Rotator(targets ...ltt.Target) ltt.Target {
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
	targets []ltt.Target
	cursor  int
	lock    sync.Mutex
}

func (rt *rotateTarget) Handle(e ltt.Event) error {
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
