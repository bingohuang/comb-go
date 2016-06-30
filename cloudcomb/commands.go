package cc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Cmd struct {
	Desc  string
	Alias string
	Func  func(args []string, opts map[string]interface{})
	Flags map[string]CmdFlag
}

type CmdFlag struct {
	Usage string
	Typ   string
}

var (
	conf   *Config
	user   *userInfo
	Driver *CcDriver

	// config file
	configFile = filepath.Join(os.Getenv("HOME"), ".cc.cfg")
)

var CmdMap = map[string]Cmd{
	"auth": {"auth in CloudComb with app key, app secret", "", Auth, nil},
	"lsco": {"list containers in CloudComb", "", LsCo, nil},
}

func Auth(args []string, opts map[string]interface{}) {
	user := &userInfo{}
	if len(args) == 2 {
		user.AppKey = args[0]
		user.AppSecret = args[1]
	} else {
		fmt.Printf("AppKey: ")
		fmt.Scanf("%s \n", &user.AppKey)
		fmt.Printf("AppSecret: ")
		fmt.Scanf("%s \n", &user.AppSecret)
	}

	driver, err := NewCcDriver(user.AppKey, user.AppSecret, 10, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Auth fail. %v", err)
		os.Exit(-1)
	}

	user.Token = driver.cc.Token
	user.ExpiresIn = driver.cc.ExpiresIn

	// save
	conf.UpdateUserInfo(user)
	conf.Save(configFile)

	fmt.Fprintf(os.Stdout, "Auth success.\n")
}

func LsCo(args []string, opts map[string]interface{}) {

}

func init() {
	if runtime.GOOS == "windows" {
		configFile = filepath.Join(os.Getenv("USERPROFILE"), ".cc.cfg")
	}
	conf = &Config{}
	conf.Load(configFile)

	user = conf.GetCurUser()
	logger := log.New(os.Stdout, "cloudcomb", 0)
	if user != nil {
		var err error
		Driver, err = NewCcDriver(user.AppKey, user.AppSecret, 10, logger)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to auth. %v\n", err)
			conf.Idx = 0
			conf.RemoveUser()
			conf.Save(configFile)
			os.Exit(-1)
		}
	}
}
