package rule

import (
	"github.com/nooize/lux"
)

func And(rules ...lwr.Rule) lwr.Rule {
	return func(e lwr.Event) bool {
		for _, rule := range rules {
			if !rule(e) {
				return false
			}

		}
		return true
	}
}

func Or(rules ...lwr.Rule) lwr.Rule {
	return func(e lwr.Event) bool {
		for _, rule := range rules {
			if rule(e) {
				return true
			}
		}
		return false
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
