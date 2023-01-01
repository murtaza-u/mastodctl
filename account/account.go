package account

import (
	"context"
	"fmt"

	"github.com/murtaza-u/mastodctl/client"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

var Cmd = &Z.Cmd{
	Name:     `account`,
	Summary:  `Get account information`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		c, err := client.New()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(
			context.Background(), client.Timeout,
		)
		defer cancel()

		acc, err := c.GetAccountCurrentUser(ctx)
		if err != nil {
			return err
		}

		out, err := yaml.Marshal(acc)
		if err != nil {
			return err
		}
		fmt.Print(string(out))

		return nil
	},
}
