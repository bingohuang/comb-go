package cc

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// container command
func container(c *cli.Context) error {
	isAll := c.Bool("a")
	isImages := c.Bool("i")
	isFlow := c.Bool("f")
	isCreate := c.Bool("c")

	// -a
	if isAll {
		result, err := driver.GetContainers()
		if err != nil {
			log.Fatalf("List all containers fail. %v", err)
		}
		fmt.Printf(result)
		return nil
	}

	// -i
	if isImages {
		result, err := driver.GetContainersImages()
		if err != nil {
			log.Fatalf("List containers' images fail. %v", err)
		}
		fmt.Printf(result)
		return nil
	}

	// args
	if len(c.Args()) == 0 {
		log.Fatalf("Container command need to specify id. See '%s auth -h'.", c.App.Name)
	}

	// -c
	if isCreate {
		jsonParams := c.Args()[0]
		id, err := driver.CreateContainer(jsonParams)
		if err != nil {
			log.Fatalf("Create container(%s) fail. %v", jsonParams, err)
		}
		fmt.Printf("Container id: %d", id)
		return nil
	}

	containerId := c.Args()[0]
	// -f
	if isFlow {
		result, err := driver.GetContainerFlow(containerId)
		if err != nil {
			log.Fatalf("Get specified container(%s)'s flow fail. %v", containerId, err)
		}
		fmt.Printf(result)
		return nil
	}

	// list container
	result, err := driver.GetContainer(containerId)
	if err != nil {
		log.Fatalf("List container(%s)'s flow fail. %v", containerId, err)
	}
	fmt.Printf(result)

	return nil
}
