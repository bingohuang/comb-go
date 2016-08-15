package cc

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// secretkey command TODO
func secretKey(c *cli.Context) error {
	if driver == nil {
		log.Fatalf("Please auth first. \n")
	}

	isAll := c.Bool("a")
	isCreate := c.Bool("c")
	isDelete := c.Bool("d")

	if isAll || len(c.Args()) == 0 {
		result, err := driver.GetSecretKeys()
		if err != nil {
			log.Fatalf("List all containers fail. %v", err)
		}
		fmt.Println(result)
		return nil
	}

	if isCreate {
		// args
		if len(c.Args()) == 0 {
			log.Fatalf("Container command need to specify key name. See '%s sk -h'.", c.App.Name)
		}
		keyName := c.Args()[0]
		id, name, err := driver.CreateSecretKey(keyName)
		if err != nil {
			log.Fatalf("Create secret key(%s) fail. %v", keyName, err)
		}
		fmt.Printf("Create Secret key Name: %s, Id: %d\n", name, id)
		return nil
	}

	if len(c.Args()) == 0 {
		log.Fatalf("Container command need to specify id. See '%s sk -h'.", c.App.Name)
	}
	secretKeyId := c.Args()[0]
	if isDelete {
		err := driver.DeleteSecretKey(secretKeyId)
		if err != nil {
			log.Fatalf("Delete specified secret key(%s) fail. %v", secretKeyId, err)
		}
		fmt.Printf("Delete secret key(%s) success.\n", secretKeyId)
		return nil
	}

	// list container
	result, err := driver.GetSecretKey(secretKeyId)
	if err != nil {
		log.Fatalf("List secret keys(%s) fail. %v", secretKeyId, err)
	}
	fmt.Println(result)

	return nil
}
