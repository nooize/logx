package rule

import (
	"github.com/nooize/lux"
)

func TagEquals(key string, v any) lux.Rule {
	return func(e lux.Event) bool {
		i := e.Tags().Value(key)
		return i == v
	}
}

func TagOneOf(key string, v ...any) lux.Rule {
	return func(e lux.Event) bool {
		i := e.Tags().Value(key)
		ok := false
		for _, v := range v {
			ok = ok || i == v
		}
		return ok
	}
}

func TagExists(key string) lux.Rule {
	return func(e lux.Event) bool {
		return e.Tags().Value(key) != nil
	}
}
