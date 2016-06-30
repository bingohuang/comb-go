package cc

import (
	"github.com/bingoHuang/cloudcomb-go-sdk/cloudcomb"
	"github.com/gosuri/uiprogress"
	"log"
)

type CcDriver struct {
	cc *cloudcomb.CloudComb

	maxConc int

	logger   *log.Logger
	progress *uiprogress.Progress
}

// New CloudComb driver
func NewCCDriver(appKey, appSecret string, conc int, logger *log.Logger) (*CcDriver, error) {
	driver := &CcDriver{
		cc:      cloudcomb.NewCC(appKey, appSecret),
		maxConc: conc,
		logger:  logger,
	}

	token, expiresIn, err := driver.cc.UserToken()
	if err != nil {
		return nil, err
	}

	driver.cc.Token = token
	driver.cc.ExpiresIn = expiresIn

	return driver, nil
}

// List containers' images
func (driver *CcDriver) ListContainersImages() (string, error) {
	result, err := driver.cc.ContainersImages()
	if err != nil {
		return "", err
	}
	return result, nil
}

// List containers
func (driver *CcDriver) ListContainers() (string, error) {
	result, err := driver.cc.Containers()
	if err != nil {
		return "", err
	}
	return result, nil
}
