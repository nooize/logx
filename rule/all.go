package rule

import (
	"github.com/nooize/lux"
	"time"
)

func Any(t time.Time, l lwr.Level, msg string, tags *lwr.Tags) bool {
	return true
}
