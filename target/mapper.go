package target

import (
	"context"
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func LevelMapper(next slog.Handler, mapping func(slog.Level) slog.Level) slog.Handler {
	trg := &levelMapperTarget{
		next:    next,
		mapping: mapping,
	}
	return trg
}

type levelMapperTarget struct {
	logx.BaseHandler
	next    slog.Handler
	mapping func(slog.Level) slog.Level
}

func (ct *levelMapperTarget) Handle(ctx context.Context, rec slog.Record) error {
	// TODO implement
	return ct.next.Handle(ctx, rec)
}
