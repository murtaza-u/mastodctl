package mastodctl

import (
	"github.com/murtaza-u/mastodctl/login"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `mastodctl`,
	Summary:   `An under-featured command line Mastodon client`,
	Usage:     `<command>`,
	Copyright: `Copyright 2022 Murtaza Udaipurwala`,
	License:   `Apache 2.0`,
	Source:    `https://github.com/murtaza-u/mastodctl`,
	Issues:    `https://github.com/murtaza-u/mastodctl/issues`,
	Commands:  []*Z.Cmd{help.Cmd, conf.Cmd, login.Cmd},
}
