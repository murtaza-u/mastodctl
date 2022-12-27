package oauth

import "strings"

// O carries the set configuration parameters and a computed scope.
// Numerous methods required for OAuth authorization flow are attached
// to it.
type O struct {
	Cfg
	scope string
}

// New instatiates a new OAuth object from the given configuration.
//
// The provided config params are validated against a set of rules.
func New(c Cfg) (*O, error) {
	err := c.validate()
	if err != nil {
		return nil, err
	}

	return &O{
		Cfg:   c,
		scope: strings.Join(c.Scopes, "+"),
	}, nil
}
