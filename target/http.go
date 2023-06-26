package target

import (
	"github.com/nooize/lux"
	"net/url"
)

func ToHttp(url url.URL) (lwr.Target, error) {
	return &writerTarget{}, nil
}
