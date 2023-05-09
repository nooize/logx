package rule

import (
	"github.com/nooize/ltt"
)

func TagEquals(key string, v any) ltt.Rule {
	return func(e ltt.Event) bool {
		i := e.Tags().Value(key)
		return i == v
	}
}

func TagOneOf(key string, v ...any) ltt.Rule {
	return func(e ltt.Event) bool {
		i := e.Tags().Value(key)
		ok := false
		for _, v := range v {
			ok = ok || i == v
		}
		return ok
	}
}

func TagExists(key string) ltt.Rule {
	return func(e ltt.Event) bool {
		return e.Tags().Value(key) != nil
	}
}