package context

import (
	"strings"

	"github.com/murtaza-u/mastodctl/config"
	"github.com/rwxrob/bonzai"
)

type comp struct{}

// NewComp instantites a new completer.
func NewComp() *comp {
	return &comp{}
}

// Complete completes all the contexts defined in the config file.
func (comp) Complete(_ bonzai.Command, args ...string) []string {
	c, err := config.New()
	if err != nil {
		return nil
	}

	contexts, err := c.GetContexts()
	if err != nil {
		return nil
	}

	var names []string
	for _, ctx := range contexts {
		names = append(names, ctx.Name)
	}

	l := len(args)
	if l == 0 {
		return names
	}

	var match []string

	for _, n := range names {
		if strings.HasPrefix(n, args[l-1]) {
			match = append(match, n)
		}
	}

	return match
}
