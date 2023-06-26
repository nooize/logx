package rule

import (
	"github.com/nooize/lux"
)

func And(rules ...lux.Rule) lux.Rule {
	return func(e lux.Event) bool {
		for _, rule := range rules {
			if !rule(e) {
				return false
			}

		}
		return true
	}
}

func Or(rules ...lux.Rule) lux.Rule {
	return func(e lux.Event) bool {
		for _, rule := range rules {
			if rule(e) {
				return true
			}
		}
		return false
	}
}

func Not(rule lux.Rule) lux.Rule {
	return func(e lux.Event) bool {
		return !rule(e)
	}
}

func True() lux.Rule {
	return func(_ lux.Event) bool {
		return true
	}
}

func False() lux.Rule {
	return func(_ lux.Event) bool {
		return false
	}
}
