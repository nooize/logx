package rule

import (
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func And(rules ...logx.Rule) logx.Rule {
	return func(e slog.Record) bool {
		for _, rule := range rules {
			if !rule(e) {
				return false
			}

		}
		return true
	}
}

func Or(rules ...logx.Rule) logx.Rule {
	return func(e slog.Record) bool {
		for _, rule := range rules {
			if rule(e) {
				return true
			}
		}
		return false
	}
}

func Not(rule logx.Rule) logx.Rule {
	return func(rec slog.Record) bool {
		return !rule(rec)
	}
}

func Bool(v bool) logx.Rule {
	return func(_ slog.Record) bool {
		return v
	}
}

func True() logx.Rule {
	return Bool(true)
}

func False() logx.Rule {
	return Bool(false)
}
