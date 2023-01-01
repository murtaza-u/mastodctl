package client

import (
	"time"

	"github.com/murtaza-u/mastodctl/config"

	"github.com/mattn/go-mastodon"
)

const Timeout = time.Second * 30

// New instantiates a new Mastodon client from the config file.
func New() (*mastodon.Client, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	ctx, err := cfg.GetContext()
	if err != nil {
		return nil, err
	}

	c := mastodon.NewClient(&mastodon.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AccessToken:  ctx.AccessToken,
		Server:       cfg.Server,
	})

	return c, nil
}
