package rule

import (
	"github.com/nooize/lux"
	"time"
)

func Any(t time.Time, l lux.Level, msg string, tags *lux.Tags) bool {
	return true
}
