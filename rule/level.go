package rule

import (
	"github.com/nooize/lux"
)

func LevelGreater(level lwr.Level) lwr.Rule {
	return func(e lwr.Event) bool {
		return level > e.Level()
	}
}

func LevelLower(level lwr.Level) lwr.Rule {
	return func(e lwr.Event) bool {
		return level < e.Level()
	}
}

func Level(levels ...lwr.Level) lwr.Rule {
	return func(e lwr.Event) bool {
		for _, v := range levels {
			if v == e.Level() {
				return true
			}
		}
		return false
	}
}
