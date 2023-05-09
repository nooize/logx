package rule

import (
	"github.com/nooize/ltt"
)

func And(rules ...ltt.Rule) ltt.Rule {
	return func(e ltt.Event) bool {
		res := true
		for _, rule := range rules {
			res = res && rule(e)
		}
		return res
	}
}

func Or(rules ...ltt.Rule) ltt.Rule {
	return func(e ltt.Event) bool {
		res := false
		for _, rule := range rules {
			res = res || rule(e)
		}
		return res
	}
}

func Not(rule ltt.Rule) ltt.Rule {
	return func(e ltt.Event) bool {
		return !rule(e)
	}
}

func True() ltt.Rule {
	return func(_ ltt.Event) bool {
		return true
	}
}

func False() ltt.Rule {
	return func(_ ltt.Event) bool {
		return false
	}
}
