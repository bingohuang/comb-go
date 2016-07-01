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
		"i": CmdFlag{"List all containers' images", "bool"},
		"a": CmdFlag{"List all containers", "bool"},
		"f": CmdFlag{"Get specified container's flow", "bool"},
	}
	ClusterFlags = map[string]CmdFlag{
		"a": CmdFlag{"List all clusters", "bool"},
	}
	RepositoryFlags = map[string]CmdFlag{
		"a": CmdFlag{"List all repositories", "bool"},
	}
	SecretkeyFlags = map[string]CmdFlag{
		"a": CmdFlag{"List all secret kyes", "bool"},
	}
	CmdMap = map[string]Cmd{
		"auth":       {"Auth in CloudComb with app key, app secret", "", Auth, nil},
		"container":  {"Container related API", "co", Container, ContainerFlags},
		"cluster":    {"Cluster related API", "cl", Cluster, ClusterFlags},
		"repository": {"Repository related API", "re", Repository, RepositoryFlags},
		"secretkey":  {"Sercet key related API", "se", Secretkey, SecretkeyFlags},
	}
)

// Auth function
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

// Container function
func Container(args []string, opts map[string]interface{}) {
	// args
	var containerId string
	if len(args) > 0 {
		containerId = args[0]
	}

	// opts
	isAll, isImages, isFlow := false, false, false
	if v, ok := opts["a"]; ok {
		if v.(bool) {
			isAll = true
		}
	}
	if v, ok := opts["i"]; ok {
		if v.(bool) {
			isImages = true
		}
	}
	if v, ok := opts["f"]; ok {
		if v.(bool) {
			isFlow = true
		}
	}

	if isAll {
		result, err := Driver.ListContainers()
		if err != nil {
			fmt.Fprintf(os.Stderr, "List containers fail. %v", err)
			os.Exit(-1)
		}
		fmt.Printf(result)
		return
	}

	if isImages {
		result, err := Driver.ListContainersImages()
		if err != nil {
			fmt.Fprintf(os.Stderr, "List containers' images fail. %v", err)
			os.Exit(-1)
		}
		fmt.Printf(result)
		return
	}

	if isFlow {
		result, err := Driver.ContainerFlow(containerId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Get specified container(%s)'s flow fail. %v", containerId, err)
			os.Exit(-1)
		}
		fmt.Printf(result)
		return
	}

	result, err := Driver.ListContainer(containerId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "List container(%s) fail. %v", containerId, err)
		os.Exit(-1)
	}
	fmt.Printf(result)
	return
}

// Cluster function TODO
func Cluster(args []string, opts map[string]interface{}) {

}

// Repository function TODO
func Repository(args []string, opts map[string]interface{}) {

}

// Secretkey function TODO
func Secretkey(args []string, opts map[string]interface{}) {

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
