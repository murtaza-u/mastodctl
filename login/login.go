package login

import (
	"fmt"

	"github.com/murtaza-u/mastodctl/config"
	"github.com/murtaza-u/mastodctl/oauth"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `login`,
	Summary:  `login to Mastodon instance`,
	Usage:    `<profile name>`,
	Commands: []*Z.Cmd{help.Cmd},
	NumArgs:  1,
	Call: func(caller *Z.Cmd, args ...string) error {
		name := args[0]
		cfg, err := config.New()
		if err != nil {
			return err
		}

		err = cfg.Validate()
		if err != nil {
			return fmt.Errorf(
				"failed to validated config file: %s", err.Error(),
			)
		}

		o, err := oauth.New(oauth.Cfg{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Server:       cfg.Server,
			Scopes:       []string{"read", "write", "follow"},
		})
		if err != nil {
			return err
		}

		code := o.AuthCode()
		tkn, err := o.Token(code)
		if err != nil {
			return fmt.Errorf(
				"failed to authorize account: %s", err.Error(),
			)
		}

		ctxs := cfg.GetContexts()
		if ctxs == nil {
			ctxs = make([]config.Ctx, 0)
		}
		ctxs = append(ctxs, config.Ctx{Name: name, AccessToken: tkn})
		cfg.Contexts = ctxs

		if cfg.Current == "" {
			cfg.Current = name
		}

		if err := cfg.Save(); err != nil {
			return err
		}

		return nil
	},
}
