package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	comb "github.com/bingoHuang/cloudcomb-go-cli/cloudcomb"
	"github.com/bingoHuang/cloudcomb-go-cli/version"
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
	app.Usage = `is a tool to manage CloudComb resources base on cloudcomb-go-sdk.`
	app.Author = "Bingo Huang"
	app.Email = "me@bingohuang.com"
	app.Version = fmt.Sprintf("%s %s/%s %s", version.VERSION, runtime.GOOS,
		runtime.GOARCH, runtime.Version())

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},

		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, info, warn, error, fatal, panic)"),
		},
	}

	// setup log-level
	app.Before = func(c *cli.Context) error {
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)

		if !c.IsSet("log-level") && !c.IsSet("l") && c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		return nil
	}

	app.Commands = make([]cli.Command, 0)
	//sort.Strings(cmds)
	for _, cmd := range cmds {
		cm, exist := comb.CmdMap[cmd]
		if exist {
			Cmd := cli.Command{
				Name:  cmd,
				Usage: cm.Desc,
				Action: func(c *cli.Context) error {
					if c.Command.FullName() != "auth" && comb.Driver == nil {
						log.Fatalln("Auth first.")
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
					cm.Func(c, opts)
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

	// run app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
