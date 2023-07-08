package console

import "io"

func Colored() TargetOption {
	return func(t *consoleHandler) {
		t.colored = true
	}
}

func WithWriter(w io.Writer) TargetOption {
	return func(t *consoleHandler) {
		if w != nil {
			t.out = w
		}
	}
}

func TagOrder(tags []string) TargetOption {
	return func(t *consoleHandler) {
		t.tags = tags
	}
}
