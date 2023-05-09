package rule

import (
	"github.com/nooize/ltt"
)

func LevelGreater(level ltt.Level) ltt.Rule {
	return func(e ltt.Event) bool {
		return level > e.Level()
	}
}

func LevelLower(level ltt.Level) ltt.Rule {
	return func(e ltt.Event) bool {
		return level < e.Level()
	}
}

func Level(levels ...ltt.Level) ltt.Rule {
	return func(e ltt.Event) bool {
		for _, v := range levels {
			if v == e.Level() {
				return true
			}
		}
		return false
	}
}
