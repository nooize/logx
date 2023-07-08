package target

import (
	"golang.org/x/exp/slog"
	"net/url"
)

func ToHttp(url url.URL) (slog.Handler, error) {
	return &writerTarget{}, nil
}
