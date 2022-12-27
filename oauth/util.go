package oauth

import (
	"net/url"

	"github.com/cli/go-gh/pkg/browser"
)

// isUrl checks if the given string is a URL.
func isUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// browse open a url in the user's default web browser.
func browse(url string) error {
	b := browser.New("", nil, nil)
	return b.Browse(url)
}
