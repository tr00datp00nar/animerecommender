package animerecommender

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/tr00datp00nar/fn"
)

var Cmd = &Z.Cmd{
	Name:        `rec`,
	Usage:       `QUERY`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_rec),
	Description: help.D(_rec),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		getRecommendation(query)
		return nil
	},
}
