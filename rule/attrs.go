package rule

import (
	"github.com/nooize/logx"
	"golang.org/x/exp/slog"
)

func AttrEquals(val slog.Attr) logx.Rule {
	return func(e slog.Record) bool {
		res := false
		e.Attrs(func(a slog.Attr) bool {
			res = a.Equal(val)
			return !res
		})
		return res
	}
}

/*
	func TagOneOf(key string, v ...any) lux.Rule {
		return func(e slog.Record) bool {
			ok := true
			e.Attrs(func(a slog.Attr) bool {
				if a.Key != key {
					return true
				}
				for _, v := range v {
					ok = ok || a.Value == v
				}
				res = a.Value.Any() == v
				return false
			})
			return res



			i := e.Tags().Value(key)
			ok := false
			for _, v := range v {
				ok = ok || i == v
			}
			return ok
		}
	}
*/
func AttrExists(key string) logx.Rule {
	return func(e slog.Record) bool {
		res := false
		e.Attrs(func(a slog.Attr) bool {
			res = a.Key == key
			return !res
		})
		return res
	}
}
