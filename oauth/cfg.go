package oauth

import (
	"errors"
	"strings"
)

// Cfg carries all the configuration parameters required to
// instantiate an OAuth instance.
type Cfg struct {
	ClientID     string
	ClientSecret string
	Server       string
	Scopes       []string
}

const RedirectURI = "urn:ietf:wg:oauth:2.0:oob"

var (
	ErrMissingClientID     = errors.New("missing client ID")
	ErrMissingClientSecret = errors.New("missing client secret")
	ErrMissingServer       = errors.New("missing server")
	ErrMissingScopes       = errors.New("missing scopes")
	ErrInvalidHttpURL      = errors.New("invalid http URL")
)

// validate validates the configuration against a set of rules.
func (c *Cfg) validate() error {
	if c.ClientID == "" {
		return ErrMissingClientID
	}

	if c.ClientSecret == "" {
		return ErrMissingClientSecret
	}

	if c.Scopes == nil || len(c.Scopes) == 0 {
		return ErrMissingScopes
	}

	if c.Server == "" {
		return ErrMissingServer
	}
	if !isUrl(c.Server) {
		return ErrInvalidHttpURL
	}
	c.Server = strings.TrimSuffix(c.Server, "/")

	return nil
}
