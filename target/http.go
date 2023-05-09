package target

import (
	"github.com/nooize/ltt"
	"net/url"
)

func ToHttp(url url.URL) (ltt.Target, error) {
	return &writerTarget{}, nil
}
