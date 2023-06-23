package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v3"
	"strings"
	"xssh/internal/xssh"
)

var AddCmd = &cli.Command{
	Name:   "add",
	Action: add,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "a",
			Usage: "别名",
		},
		&cli.StringFlag{
			Name:  "c",
			Usage: "root@xxxx 配置",
		},
		&cli.BoolFlag{
			Name:  "f",
			Usage: "强行覆盖",
		},
		&cli.StringFlag{
			Name:  "i",
			Usage: "ssh 配置文件 默认 ~/.ssh/id_rsa",
			Value: "~/.ssh/id_rsa",
		},
		&cli.StringFlag{
			Name:  "p",
			Usage: "ssh  端口，默认 22",
			Value: "22",
		},
	},
}

func add(ctx *cli.Context) error {
	overwrite := ctx.Bool("f")
	alias := ctx.String("a")
	userHost := ctx.String("c")
	parts := strings.Split(userHost, "@")
	if len(parts) != 2 {
		return errors.New("-c option must be: xxx@somehost.com")
	}
	port := ctx.String("p")
	i := ctx.String("i")
	mdl := xssh.XSSHModel{
		Port:     port,
		SshKey:   i,
		Username: parts[0],
		Host:     parts[1],
		Alias:    alias,
	}

	fmt.Println(mdl)
	return xssh.Add(&mdl, overwrite)
}
