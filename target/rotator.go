package target

import (
	"context"
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
	"sync"
)

func Rotator(targets ...slog.Handler) slog.Handler {
	trg := &rotateHandler{
		targets: targets,
		cursor:  0,
		lock:    sync.Mutex{},
	}
	if len(targets) == 0 {
		trg.cursor = -1
	}
	return trg
}

type rotateHandler struct {
	logx.BaseHandler
	targets []slog.Handler
	cursor  int
	lock    sync.Mutex
}

func (rt *rotateHandler) Handle(ctx context.Context, rec slog.Record) error {
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
	return target.Handle(ctx, rec)
}
