package main

import (
	cc "./cloudcomb"
	"./version"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

var cmds = []string{
	"auth", "container", "cluster", "repository", "secretkey",
}

func main() {
	app := cli.NewApp()
	app.Name = "comb"
	app.Usage = `is a tool for manage resources in [CloudComb](http://c.163.com)
		    base on [cloudcomb-go-sdk](https://github.com/bingoHuang/cloudcomb-go-cli)`
	app.Author = "Bingo Huang"
	app.Email = "bingo@xbing.me"
	app.Version = fmt.Sprintf("%s %s/%s %s", version.VERSION, runtime.GOOS,
		runtime.GOARCH, runtime.Version())
	app.Commands = make([]cli.Command, 0)

	//sort.Strings(cmds)

	for _, cmd := range cmds {
		cm, exist := cc.CmdMap[cmd]
		if exist {
			Cmd := cli.Command{
				Name:  cmd,
				Usage: cm.Desc,
				Action: func(c *cli.Context) error {
					if c.Command.FullName() != "auth" && cc.Driver == nil {
						fmt.Println("Auth first.")
						os.Exit(-1)
					}
					opts := make(map[string]interface{})
					for k, v := range cm.Flags {
						if c.IsSet(k) {
							switch v.Typ {
							case "bool":
								opts[k] = c.Bool(k)
							case "string":
								opts[k] = c.String(k)
							case "int":
								opts[k] = c.Int(k)
							}

						}
					}
					cm.Func(c.Args(), opts)
					return nil
				},
			}
			if cm.Alias != "" {
				Cmd.Aliases = []string{cm.Alias}
			}
			if cm.Flags != nil {
				Cmd.Flags = []cli.Flag{}
				for k, v := range cm.Flags {
					var flag cli.Flag
					switch v.Typ {
					case "bool":
						flag = cli.BoolFlag{Name: k, Usage: v.Usage}
					case "string":
						flag = cli.StringFlag{Name: k, Usage: v.Usage}
					case "int":
						flag = cli.IntFlag{Name: k, Usage: v.Usage}
					}
					Cmd.Flags = append(Cmd.Flags, flag)
				}
			}
			app.Commands = append(app.Commands, Cmd)
		}
	}

	app.Run(os.Args)
}
