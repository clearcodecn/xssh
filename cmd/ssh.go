package cmd

import (
	"github.com/clearcodecn/xssh/internal/xssh"
	"github.com/urfave/cli/v3"
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
