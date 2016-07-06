package cc

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	d "github.com/bingoHuang/cloudcomb-go-cli/driver"
	"github.com/codegangsta/cli"
	"os"
)

// auth command
func auth(c *cli.Context) error {
	//user := &userInfo{}
	if len(c.Args()) == 2 {
		user.AppKey = c.Args()[0]
		user.AppSecret = c.Args()[1]
	} else if len(c.Args()) == 0 {
		fmt.Printf("AppKey: ")
		fmt.Scanf("%s \n", &user.AppKey)
		fmt.Printf("AppSecret: ")
		fmt.Scanf("%s \n", &user.AppSecret)
	} else {
		log.Fatalf("Auth command takes exact two argument. See '%s auth -h'.", c.App.Name)
	}

	var err error
	driver, err = d.NewCCDriver(user.AppKey, user.AppSecret, 10)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Auth fail. %v", err)
		os.Exit(-1)
	}

	user.Token = driver.Cc.Token
	user.ExpiresIn = driver.Cc.ExpiresIn

	// save
	conf.UpdateUserInfo(user)
	conf.Save(configFile)

	fmt.Printf("Auth success.\n")
	return nil
}
