package oauth

import (
	"net/url"

	"github.com/cli/go-gh/pkg/browser"
)

func isUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func browse(url string) error {
	b := browser.New("", nil, nil)
	return b.Browse(url)
}
