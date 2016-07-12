package cc

import (
	log "github.com/Sirupsen/logrus"
	c "github.com/bingoHuang/comb/config"
	d "github.com/bingoHuang/comb/driver"
	"github.com/codegangsta/cli"
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
			Name:   "auth",
			Usage:  "Auth in CloudComb with app key, app secret",
			Action: auth,
		},
		{
			Name:    "container",
			Usage:   "Operate containers in CloudComb",
			Aliases: []string{"co"},
			Action:  container,
			Flags:   containerFlags,
		},
		{
			Name:    "cluster",
			Usage:   "Operate clusters in CloudComb",
			Aliases: []string{"cu"},
			Action:  cluster,
			Flags:   clusterFlags,
		},
		{
			Name:    "repositry",
			Usage:   "Operate repositrys in CloudComb",
			Aliases: []string{"re"},
			Action:  repository,
			Flags:   repositoryFlags,
		},
		{
			Name:    "secretkey",
			Usage:   "Operate secret keys in CloudComb",
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
			conf.Idx = 0
			conf.RemoveUser()
			conf.Save(configFile)
			log.Fatalf("Failed to auth. %v\n", err)
		}
	} else {
		user = new(c.UserInfo)
	}
}
