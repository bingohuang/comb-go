package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Cmd struct {
	Desc  string
	Alias string
	Func  func(args []string, opts map[string]interface{})
	Flags map[string]CmdFlag
}

type CmdFlag struct {
	usage string
	typ   string
}

var (
	conf   *Config
	//user   *userInfo
	driver *CcDriver

	// config file
	configFile = filepath.Join(os.Getenv("HOME"), ".cc.cfg")
)

var CmdMap = map[string]Cmd{
	"auth": {"get auth in CloudComb with app key, app secret", "", Auth, nil},
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
