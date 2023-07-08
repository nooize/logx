package target

import (
	"context"
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func Сondition(first, second slog.Handler, rule logx.Rule) slog.Handler {
	return &conditionTarget{
		first:  first,
		second: second,
		rule:   rule,
	}
}

type conditionTarget struct {
	logx.BaseHandler
	first  slog.Handler
	second slog.Handler
	rule   logx.Rule
}

func (ct *conditionTarget) Handle(ctx context.Context, rec slog.Record) error {
	if ct.rule(rec) {
		return ct.first.Handle(ctx, rec)
	}
	return ct.second.Handle(ctx, rec)
}
