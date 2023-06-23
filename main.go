package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"log"
	"os"
	"xssh/cmd"
)

var (
	rootCmd = cli.Command{
		Name:  "xssh",
		Usage: "xssh 管理多个 ssh 服务器\n 使用 xssh [alias] 登录服务器",
	}
)

func init() {
	rootCmd.Commands = []*cli.Command{
		cmd.AddCmd,
		cmd.ListCmd,
		cmd.SSHCmd,
		cmd.DelCmd,
	}
}

func main() {
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Println(err)
	}
}
