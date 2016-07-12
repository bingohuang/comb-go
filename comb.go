package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	combCli "github.com/bingoHuang/comb/cli"
	"github.com/bingoHuang/comb/version"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
)

func main() {
	app := cli.NewApp()
	app.Name = "comb"
	app.Usage = `is a tool to manage CloudComb resources base on cloudcomb-go-sdk.`
	app.Version = fmt.Sprintf("%s %s/%s %s", version.VERSION, runtime.GOOS,
		runtime.GOARCH, runtime.Version())

	app.Author = "Bingo Huang"
	app.Email = "me@bingohuang.com"

	// setup global flags
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

	// setup log level
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

	// setup commands
	app.Commands = combCli.Commands

	// run app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
