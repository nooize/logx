package rule

import (
	"github.com/nooize/lwr"
	"time"
)

func Any(t time.Time, l lwr.Level, msg string, tags *lwr.Tags) bool {
	return true
}
