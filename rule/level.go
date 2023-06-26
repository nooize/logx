package rule

import (
	"github.com/nooize/lux"
)

func LevelGreater(level lux.Level) lux.Rule {
	return func(e lux.Event) bool {
		return level > e.Level()
	}
}

func LevelLower(level lux.Level) lux.Rule {
	return func(e lux.Event) bool {
		return level < e.Level()
	}
}

func Level(levels ...lux.Level) lux.Rule {
	return func(e lux.Event) bool {
		for _, v := range levels {
			if v == e.Level() {
				return true
			}
		}
		return false
	}
}
