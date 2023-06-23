package cmd

import (
	"github.com/urfave/cli/v3"
	"xssh/internal/xssh"
)

var ListCmd = &cli.Command{
	Name:   "list",
	Action: list,
}

func list(ctx *cli.Context) error {
	xssh.List(ctx.String("substr"))
	return nil
}
