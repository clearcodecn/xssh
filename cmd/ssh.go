package cmd

import (
	"github.com/urfave/cli/v3"
	"xssh/internal/xssh"
)

var SSHCmd = &cli.Command{
	Name: "ssh",
	Action: func(ctx *cli.Context) error {
		return xssh.Ssh()
	},
}

var DelCmd = &cli.Command{
	Name: "del",
	Action: func(ctx *cli.Context) error {
		return xssh.Del()
	},
}
