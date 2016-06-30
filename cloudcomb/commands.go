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

var (
	ContainerFlags = map[string]CmdFlag{
		"a": CmdFlag{"list all containers", "string"},
	}
)

var CmdMap = map[string]Cmd{
	"auth": {"auth in CloudComb with app key, app secret", "", Auth, nil},
	"lsci": {"list all containers' images", "", LsCI, nil},
	"lscs": {"list all containers' info", "", LsCs, nil},
	"lsco": {"list specified container's info with id or name", "", LsCo, nil},
	"flow": {"Get specified container's flow with id or name ", "", LsCo, nil},

	"container": {"Container related API", "co", LsCI, ContainerFlags},
	"cluster": {"Cluster related API", "cl", LsCI, nil},
	"repository": {"Repository related API", "re", LsCI, nil},
	"secretkey": {"Sercet key related API", "sk", LsCI, nil},
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

	driver, err := NewCCDriver(user.AppKey, user.AppSecret, 10, nil)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Auth fail. %v", err)
		os.Exit(-1)
	}

	user.Token = driver.cc.Token
	user.ExpiresIn = driver.cc.ExpiresIn

	// save
	conf.UpdateUserInfo(user)
	conf.Save(configFile)

	fmt.Printf("Auth success.\n")
}

func LsCI(args []string, opts map[string]interface{}) {
	result, err := Driver.ListContainersImages()
	if err != nil {
		fmt.Fprintf(os.Stderr, "List containers' images fail. %v", err)
		os.Exit(-1)
	}
	fmt.Printf(result)
}

func LsCs(args []string, opts map[string]interface{}) {
	result, err := Driver.ListContainers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "List containers fail. %v", err)
		os.Exit(-1)
	}
	fmt.Printf(result)
}

func LsCo(args []string, opts map[string]interface{}) {
	result, err := Driver.ListContainers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "List containers fail. %v", err)
		os.Exit(-1)
	}
	fmt.Printf(result)
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
		Driver, err = NewCCDriver(user.AppKey, user.AppSecret, 10, logger)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to auth. %v\n", err)
			conf.Idx = 0
			conf.RemoveUser()
			conf.Save(configFile)
			os.Exit(-1)
		}
	}
}
