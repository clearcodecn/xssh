package cmd

import (
	"github.com/clearcodecn/xssh/internal/xssh"
	"github.com/urfave/cli/v3"
)

var ListCmd = &cli.Command{
	Name:   "list",
	Action: list,
}

func list(ctx *cli.Context) error {
	xssh.List(ctx.String("substr"))
	return nil
}
