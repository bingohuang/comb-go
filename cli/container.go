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

	// -a
	if isAll {
		result, err := driver.ListContainers()
		if err != nil {
			log.Fatalf("List all containers fail. %v", err)
		}
		fmt.Printf(result)
		return nil
	}

	// -i
	if isImages {
		result, err := driver.ListContainersImages()
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

	containerId := c.Args()[0]

	// -f
	if isFlow {
		result, err := driver.ContainerFlow(containerId)
		if err != nil {
			log.Fatalf("Get specified container(%s)'s flow fail. %v", containerId, err)
		}
		fmt.Printf(result)
		return nil
	}

	// list container
	result, err := driver.ListContainer(containerId)
	if err != nil {
		log.Fatalf("List container(%s)'s flow fail. %v", containerId, err)
	}
	fmt.Printf(result)

	return nil
}
