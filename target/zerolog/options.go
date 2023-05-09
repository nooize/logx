package zerolog

import (
	"github.com/rs/zerolog"
)

type Option func(t *zerologTarget)

func WithLogger(logger zerolog.Logger) Option {
	return func(t *zerologTarget) {
		t.logger = logger
	}
}
