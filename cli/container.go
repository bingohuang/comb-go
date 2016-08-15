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
		fmt.Println(result)
		return nil
	}

	// -a
	if isAll || len(c.Args()) == 0 {
		result, err := driver.GetContainers()
		if err != nil {
			log.Fatalf("List all containers fail. %v", err)
		}
		fmt.Println(result)
		return nil
	}

	// args
	if len(c.Args()) == 0 {
		log.Fatalf("Container command need to specify param. See '%s co -h'.", c.App.Name)
	}

	// -c
	if isCreate {
		jsonParams := c.Args()[0]
		id, err := driver.CreateContainer(jsonParams)
		if err != nil {
			log.Fatalf("Create container(%s) fail. %v", jsonParams, err)
		}
		fmt.Printf("Create Container(Id: %d) success.\n", id)
		return nil
	}
	containerId := c.Args()[0]
	// -f
	if isFlow {
		result, err := driver.GetContainerFlow(containerId)
		if err != nil {
			log.Fatalf("Get specified container(%s) flow fail. %v", containerId, err)
		}
		fmt.Println(result)
		return nil
	}
	// -r
	if isRestart {
		err := driver.RestartContainer(containerId)
		if err != nil {
			log.Fatalf("Restart specified container(%s) fail. %v", containerId, err)
		}
		fmt.Printf("Restarting container(%s)...\n", containerId)
		return nil
	}
	// -d
	if isDelete {
		err := driver.DeleteContainer(containerId)
		if err != nil {
			log.Fatalf("Delete specified container(%s) fail. %v", containerId, err)
		}
		fmt.Printf("Delete container(%s) success.\n", containerId)
		return nil
	}

	// -t|-u args
	if (isTag || isUpdate) && len(c.Args()) < 2 {
		log.Fatalf("Container command need to specify id and params. See '%s co -h'.", c.App.Name)
	}
	// -t
	if isTag {
		jsonParams := c.Args()[1]
		result, err := driver.TagContainer(containerId, jsonParams)
		if err != nil {
			log.Fatalf("Tag specified container(%s) with content(%s) fail. %v", containerId, jsonParams, err)
		}
		fmt.Println(result)
		return nil
	}
	// -u
	if isUpdate {
		jsonParams := c.Args()[1]
		err := driver.UpdateContainer(containerId, jsonParams)
		if err != nil {
			log.Fatalf("Update specified container(%s) with content(%s) fail. %v", containerId, jsonParams, err)
		}
		fmt.Printf("Update container(%s) success.\n", containerId)
		return nil
	}

	// list container
	result, err := driver.GetContainer(containerId)
	if err != nil {
		log.Fatalf("List container(%s) fail. %v", containerId, err)
	}
	fmt.Println(result)

	return nil
}
