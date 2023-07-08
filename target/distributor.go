package target

import (
	"context"
	"errors"
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func Distributor(targets ...slog.Handler) slog.Handler {
	return &distributeHandler{
		targets: targets,
	}
}

type distributeHandler struct {
	logx.BaseHandler
	targets []slog.Handler
}

func (dh *distributeHandler) Handle(ctx context.Context, rec slog.Record) error {
	count := len(dh.targets)
	if count == 0 {
		return nil
	}
	results := make(chan error, count)
	defer close(results)
	for _, trg := range dh.targets {
		go func(trg slog.Handler) {
			results <- trg.Handle(ctx, rec)
		}(trg)
	}
	var out error
	for i := 0; i < count; i++ {
		if err := <-results; err != nil {
			errors.Join(out, err)
		}
	}
	return out
}
