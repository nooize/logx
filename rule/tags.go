package rule

import (
	"github.com/nooize/lwr"
)

func TagEquals(key string, v any) lwr.Rule {
	return func(e lwr.Event) bool {
		i := e.Tags().Value(key)
		return i == v
	}
}

func TagOneOf(key string, v ...any) lwr.Rule {
	return func(e lwr.Event) bool {
		i := e.Tags().Value(key)
		ok := false
		for _, v := range v {
			ok = ok || i == v
		}
		return ok
	}
}

func TagExists(key string) lwr.Rule {
	return func(e lwr.Event) bool {
		return e.Tags().Value(key) != nil
	}
}
