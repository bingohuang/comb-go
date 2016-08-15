package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

var (
	appKey    string
	appSecret string
)

func init() {
	appKey = os.Getenv("CC_APPKEY")
	appSecret = os.Getenv("CC_APPSECRET")
}

/*
	Test cases
*/
func TestAuth(t *testing.T) {
	res, err := combCli("auth", appKey, appSecret)

	check(t, err == nil, "failed to auth")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestGetContainersImages(t *testing.T) {
	res, err := combCli("co", "-i")

	check(t, err == nil, "failed to get all containers images")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestGetContainers(t *testing.T) {
	res, err := combCli("co")

	check(t, err == nil, "failed to get all containers by `comb co`")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))

	res, err = combCli("co", "-a")

	check(t, err == nil, "failed to get all containers by `comb co -a`")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestCreateContainer(t *testing.T) {
	params := `{
				"image_type": %d,
				"image_id": %d,
				"name": "%s",
				"desc": "%s",
				"ssh_key_ids": %s,
				"env_var": {
					"key1": "value1",
    				"key2": "value2"
				},
				"charge_type": 2,
				"spec_id": 1,
				"use_public_network": 1,
				"network_charge_type": 1,
				"bandwidth": 1
			  }`
	params = purifyParams(fmt.Sprintf(params, 1, 20835, "openapitest", "cloudcomb open api test container", "[]"))

	res, err := combCli("co", "-c", params)

	check(t, err == nil, "failed to create container", err)

	resStr := string(res)
	// substring the container id from "Container id: %d\n"
	containerId := resStr[strings.LastIndex(resStr, ":")+2 : len(resStr)-1]

	os.Setenv("CC_CONTAINER_ID", containerId)
	fmt.Printf("CC_CONTAINER_ID=%s\n", os.Getenv("CC_CONTAINER_ID"))

	// waiting 60s for finish creating the container
	time.Sleep(time.Duration(60 * time.Second))
}

func TestGetContainer(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	res, err := combCli("co", containerId)

	check(t, err == nil, fmt.Sprintf("failed to get the container(id=%s)", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestGetContainerFlow(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	res, err := combCli("co", "-f", containerId)

	check(t, err == nil, fmt.Sprintf("failed to get the container(id=%s) flow", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestTagContainer(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	params := `{
	  "repo_name": "%s",
	  "tag": "%s"
	}`
	params = purifyParams(fmt.Sprintf(params, "openapi", time.Now().Format("20060102150405")))
	//params = fmt.Sprintf("%q", params)

	res, err := combCli("co", "-t", containerId, params)

	check(t, err == nil, fmt.Sprintf("failed to tag the container(id=%s)", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestUpdateContainer(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	params := `{
	  "charge_type":%d,
	  "desc": "%s",
	  "network_charge_type":%d,
	  "bandwidth":%d
	}`
	params = purifyParams(fmt.Sprintf(params, 2, "Modify description", 1, 2))
	//params = fmt.Sprintf("%q", params)

	res, err := combCli("co", "-u", containerId, params)

	check(t, err == nil, fmt.Sprintf("failed to update the container(id=%s)", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(10 * time.Second))
}

func TestRestartContainer(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	res, err := combCli("co", "-r", containerId)

	check(t, err == nil, fmt.Sprintf("failed to restart the container(id=%s)", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(30 * time.Second))
}

func TestDeleteContainer(t *testing.T) {
	containerId := os.Getenv("CC_CONTAINER_ID")
	check(t, containerId != "", "failed to get the container id")

	res, err := combCli("co", "-d", containerId)

	check(t, err == nil, fmt.Sprintf("failed to delete the container(id=%s)", containerId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestGetSecretKeys(t *testing.T) {
	res, err := combCli("sk")

	check(t, err == nil, "failed to get all secret keys by `comb sk`")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))

	res, err = combCli("sk", "-a")

	check(t, err == nil, "failed to get all secret keys  by `comb sk -a`")
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestCreateSecretKey(t *testing.T) {
	keyName := fmt.Sprint("test", "-", time.Now().Unix())
	res, err := combCli("sk", "-c", keyName)

	check(t, err == nil, "failed to create secret key", err)

	resStr := string(res)
	// substring the container id from "Container id: %d\n"
	secretKeyId := resStr[strings.LastIndex(resStr, ":")+2 : len(resStr)-1]

	os.Setenv("CC_SECRETKEY_ID", secretKeyId)
	fmt.Printf("CC_SECRETKEY_ID=%s\n", os.Getenv("CC_SECRETKEY_ID"))

	// waiting 60s for finish creating the container
	time.Sleep(time.Duration(3 * time.Second))
}

func TestGetSecretKey(t *testing.T) {
	secretKeyId := os.Getenv("CC_SECRETKEY_ID")
	check(t, secretKeyId != "", "failed to get the secret key id")

	res, err := combCli("sk", secretKeyId)

	check(t, err == nil, fmt.Sprintf("failed to get the secret key(id=%s)", secretKeyId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

func TestDeleteSecretKey(t *testing.T) {
	secretKeyId := os.Getenv("CC_SECRETKEY_ID")
	check(t, secretKeyId != "", "failed to get the secret key id")

	res, err := combCli("sk", "-d", secretKeyId)

	check(t, err == nil, fmt.Sprintf("failed to delete the secret key(id=%s)", secretKeyId), err)
	fmt.Printf("%s", string(res))
	time.Sleep(time.Duration(time.Second))
}

/*
	Functions for test
*/
func combCli(cmd string, args ...string) ([]byte, error) {
	args = append([]string{cmd}, args...)
	fmt.Printf("args=%s\n", args)
	return exec.Command("./comb", args...).Output()
}

func check(t *testing.T, cond bool, arg0 string, args ...interface{}) {
	if !cond {
		if !strings.HasSuffix(arg0, "\n") {
			arg0 += "\n"
		}
		t.Errorf(arg0, args...)
		if t != nil {
			t.FailNow()
		}
	}
}

// purify params by "\n", "\t"
func purifyParams(params string) string {
	params = strings.Replace(params, "\n", "", -1)
	params = strings.Replace(params, "\t", "", -1)
	return params
}
