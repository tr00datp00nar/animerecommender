package animerecommender

import (
	"net/url"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
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
		query := rawUrlEncode(a)
		getRecommendation(query)
		return nil
	},
}

// rawUrlEncode takes any string and conactenates each word with "%20"
// as the delimeter. Used when encoding a URL.
func rawUrlEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
