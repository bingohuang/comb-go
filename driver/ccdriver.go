package driver

import (
	"github.com/bingoHuang/cloudcomb-go-sdk/cloudcomb"
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

	token, expiresIn, err := driver.Cc.UserToken()
	if err != nil {
		return nil, err
	}

	driver.Cc.Token = token
	driver.Cc.ExpiresIn = expiresIn

	return driver, nil
}

// List containers' images
func (driver *CcDriver) ListContainersImages() (string, error) {
	result, err := driver.Cc.GetContainersImages()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List containers
func (driver *CcDriver) ListContainers() (string, error) {
	result, err := driver.Cc.GetContainers()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List specified container
func (driver *CcDriver) ListContainer(id string) (string, error) {
	result, err := driver.Cc.GetContainer(id)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Get specified container's flow
func (driver *CcDriver) ContainerFlow(id string) (string, error) {
	result, err := driver.Cc.GetContainerFlow(id)
	if err != nil {
		return "", err
	}
	return result, nil
}
