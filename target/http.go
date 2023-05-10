package target

import (
	"github.com/nooize/lwr"
	"net/url"
)

func ToHttp(url url.URL) (lwr.Target, error) {
	return &writerTarget{}, nil
}
