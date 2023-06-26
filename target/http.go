package target

import (
	"github.com/nooize/lux"
	"net/url"
)

func ToHttp(url url.URL) (lux.Target, error) {
	return &writerTarget{}, nil
}
