package rule

import (
	"github.com/nooize/lwr"
)

func And(rules ...lwr.Rule) lwr.Rule {
	return func(e lwr.Event) bool {
		res := true
		for _, rule := range rules {
			res = res && rule(e)
		}
		return res
	}
}

func Or(rules ...lwr.Rule) lwr.Rule {
	return func(e lwr.Event) bool {
		res := false
		for _, rule := range rules {
			res = res || rule(e)
		}
		return res
	}
}

func Not(rule lwr.Rule) lwr.Rule {
	return func(e lwr.Event) bool {
		return !rule(e)
	}
}

func True() lwr.Rule {
	return func(_ lwr.Event) bool {
		return true
	}
}

func False() lwr.Rule {
	return func(_ lwr.Event) bool {
		return false
	}
}
