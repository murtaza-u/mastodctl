package context

import (
	"fmt"

	"github.com/murtaza-u/mastodctl/config"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `context`,
	Usage:    `command`,
	Summary:  `manage mastodctl contexts`,
	Commands: []*Z.Cmd{help.Cmd, listCmd, setCmd, deleteCmd},
}

var listCmd = &Z.Cmd{
	Name:    `list`,
	Summary: `list available contexts`,
	Comp:    NewComp(),
	Call: func(caller *Z.Cmd, args ...string) error {
		cfg, err := config.New()
		if err != nil {
			return err
		}

		contexts, err := cfg.GetContexts()
		if err != nil {
			return nil
		}

		for _, ctx := range contexts {
			if cfg.Current == ctx.Name {
				ctx.Name += "(*)"
			}
			fmt.Println(ctx.Name)
		}

		return nil
	},
}

var setCmd = &Z.Cmd{
	Name:    `set`,
	Summary: `set current context`,
	Usage:   `context-name`,
	NumArgs: 1,
	Comp:    NewComp(),
	Call: func(caller *Z.Cmd, args ...string) error {
		ctx := args[0]

		cfg, err := config.New()
		if err != nil {
			return err
		}

		err = cfg.SetContext(ctx)
		if err != nil {
			return err
		}

		return cfg.Save()
	},
}

var deleteCmd = &Z.Cmd{
	Name:    `delete`,
	Usage:   `context-name`,
	Summary: `delete a context`,
	Comp:    NewComp(),
	Call: func(caller *Z.Cmd, args ...string) error {
		ctx := args[0]

		cfg, err := config.New()
		if err != nil {
			return err
		}

		err = cfg.RemoveContext(ctx)
		if err != nil {
			return err
		}

		if ctx == cfg.Current {
			cfg.Current = ""
		}

		return cfg.Save()
	},
}
