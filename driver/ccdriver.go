package driver

import (
	"fmt"
	cloudcomb "github.com/bingohuang/cloudcomb-go-sdk"
	"github.com/gosuri/uiprogress"
)

type CcDriver struct {
	Cc *cloudcomb.CloudComb

	maxConc int

	progress *uiprogress.Progress
}

// New CloudComb driver
func NewCCDriver(appKey, appSecret string, conc int) (*CcDriver, error) {
	driver := &CcDriver{
		Cc:      cloudcomb.NewCC(appKey, appSecret),
		maxConc: conc,
	}

	token, err := driver.Cc.UserToken()
	if err != nil {
		return nil, err
	}

	driver.Cc.Token = token

	return driver, nil
}

// List containers' images
func (driver *CcDriver) GetContainersImages() (string, error) {
	result, err := driver.Cc.GetContainersImages()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List containers
func (driver *CcDriver) GetContainers() (string, error) {
	result, err := driver.Cc.GetContainers()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List specified container
func (driver *CcDriver) GetContainer(id string) (string, error) {
	result, err := driver.Cc.GetContainer(id)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Get specified container's flow
func (driver *CcDriver) GetContainerFlow(id string) (string, error) {
	result, err := driver.Cc.GetContainerFlow(id)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Create container
func (driver *CcDriver) CreateContainer(params string) (uint, error) {
	result, err := driver.Cc.CreateContainer(params)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Update container
func (driver *CcDriver) UpdateContainer(id string, params string) error {
	err := driver.Cc.UpdateContainer(id, params)
	if err != nil {
		return err
	}
	return nil
}

// Restart container
func (driver *CcDriver) RestartContainer(id string) error {
	err := driver.Cc.RestartContainer(id)
	if err != nil {
		return err
	}
	return nil
}

// Tag container
func (driver *CcDriver) TagContainer(id string, params string) (string, error) {
	result, err := driver.Cc.TagContainer(id, params)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Restart container
func (driver *CcDriver) DeleteContainer(id string) error {
	err := driver.Cc.DeleteContainer(id)
	if err != nil {
		return err
	}
	return nil
}

// List containers
func (driver *CcDriver) GetSecretKeys() (string, error) {
	result, err := driver.Cc.GetSecretKeys()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List specified container
func (driver *CcDriver) GetSecretKey(id string) (string, error) {
	result, err := driver.Cc.GetSecretKey(id)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Create container
func (driver *CcDriver) CreateSecretKey(name string) (uint, string, error) {
	params := `{
				"key_name": "%s"
			}`
	params = fmt.Sprintf(params, name)
	id, name, err := driver.Cc.CreateSecretKey(params)
	if err != nil {
		return 0, "", err
	}
	return id, name, nil
}

// Restart container
func (driver *CcDriver) DeleteSecretKey(id string) error {
	err := driver.Cc.DeleteSecretKey(id)
	if err != nil {
		return err
	}
	return nil
}
