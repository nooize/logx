package rule

import (
	"github.com/nooize/ltt"
	"time"
)

func Any(t time.Time, l ltt.Level, msg string, tags *ltt.Tags) bool {
	return true
}
