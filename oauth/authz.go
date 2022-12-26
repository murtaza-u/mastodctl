package oauth

import "strings"

type O struct {
	Cfg
	scope string
}

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
