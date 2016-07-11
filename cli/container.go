package cc

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

// container command
func container(c *cli.Context) error {
	if driver == nil {
		log.Fatalf("Please auth first. \n")
	}
	isImages := c.Bool("i")
	isAll := c.Bool("a")
	isFlow := c.Bool("f")
	isCreate := c.Bool("c")
	isUpdate := c.Bool("u")
	isRestart := c.Bool("r")
	isTag := c.Bool("t")
	isDelete := c.Bool("d")

	// -i
	if isImages {
		result, err := driver.GetContainersImages()
		if err != nil {
			log.Fatalf("List containers images fail. %v", err)
		}
		fmt.Printf(result)
		return nil
	}

	// -a
	if isAll {
		result, err := driver.GetContainers()
		if err != nil {
			log.Fatalf("List all containers fail. %v", err)
		}
		fmt.Printf(result)
		return nil
	}

	// args
	if len(c.Args()) == 0 {
		log.Fatalf("Container command need to specify id. See '%s co -h'.", c.App.Name)
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
			log.Fatalf("Get specified container(%s) flow fail. %v", containerId, err)
		}
		fmt.Printf(result)
		return nil
	}
	// -u
	if isUpdate {
		jsonParams := c.Args()[1]
		err := driver.UpdateContainer(containerId, jsonParams)
		if err != nil {
			log.Fatalf("Update specified container(%s) with content(%s) fail. %v", containerId, jsonParams, err)
		}
		return nil
	}
	// -r
	if isRestart {
		err := driver.RestartContainer(containerId)
		if err != nil {
			log.Fatalf("Restart specified container(%s) fail. %v", containerId, err)
		}
		return nil
	}
	// -t
	if isTag {
		jsonParams := c.Args()[1]
		result, err := driver.TagContainer(containerId, jsonParams)
		if err != nil {
			log.Fatalf("Tag specified container(%s) with content(%s) fail. %v", containerId, jsonParams, err)
		}
		fmt.Printf(result)
		return nil
	}
	// -d
	if isDelete {
		err := driver.DeleteContainer(containerId)
		if err != nil {
			log.Fatalf("Delete specified container(%s) fail. %v", containerId, err)
		}
		return nil
	}

	// list container
	result, err := driver.GetContainer(containerId)
	if err != nil {
		log.Fatalf("List container(%s) flow fail. %v", containerId, err)
	}
	fmt.Printf(result)

	return nil
}
