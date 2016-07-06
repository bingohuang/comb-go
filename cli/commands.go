package cc

import (
	"fmt"
	"github.com/codegangsta/cli"
	c "github.com/bingoHuang/cloudcomb-go-cli/config"
	d "github.com/bingoHuang/cloudcomb-go-cli/driver"
	"os"
	"path/filepath"
	"runtime"
)

var (
	conf   *c.Config
	user   *c.UserInfo
	driver *d.CcDriver

	// config file
	configFile = filepath.Join(os.Getenv("HOME"), ".comb.cfg")
)

var (
	Commands = []cli.Command{
		{
			Name:    "auth",
			Usage:   "Auth in CloudComb with app key, app secret",
			Action:  auth,
		},
		{
			Name:    "container",
			Usage:   "Operate container",
			Aliases: []string{"co"},
			Action:  container,
			Flags:   containerFlags,
		},
		{
			Name:    "cluster",
			Usage:   "Operate cluster",
			Aliases: []string{"cu"},
			Action:  cluster,
			Flags:   clusterFlags,
		},
		{
			Name:    "repositry",
			Usage:   "Operate repositry",
			Aliases: []string{"re"},
			Action:  repository,
			Flags:   repositoryFlags,
		},
		{
			Name:    "secretkey",
			Usage:   "Operate secret key",
			Aliases: []string{"sk"},
			Action:  secretKey,
			Flags:   secretkeyFlags,
		},
	}
)

func init() {
	if runtime.GOOS == "windows" {
		configFile = filepath.Join(os.Getenv("USERPROFILE"), ".comb.cfg")
	}
	conf = &c.Config{}
	conf.Load(configFile)

	user = conf.GetCurUser()
	if user != nil {
		var err error
		driver, err = d.NewCCDriver(user.AppKey, user.AppSecret, 10)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to auth. %v\n", err)
			conf.Idx = 0
			conf.RemoveUser()
			conf.Save(configFile)
			os.Exit(-1)
		}
	}
}
