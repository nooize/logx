package rule

import (
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func LevelGreater(level slog.Level) logx.Rule {
	return func(e slog.Record) bool {
		return e.Level > level
	}
}

func LevelLower(level slog.Level) logx.Rule {
	return func(e slog.Record) bool {
		return e.Level < level
	}
}

func LevelIn(levels ...slog.Level) logx.Rule {
	return func(e slog.Record) bool {
		for _, v := range levels {
			if v == e.Level {
				return true
			}
		}
		return false
	}
}

func LevelNotIn(levels ...slog.Level) logx.Rule {
	return func(e slog.Record) bool {
		for _, v := range levels {
			if v == e.Level {
				return false
			}
		}
		return true
	}
}
