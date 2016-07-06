package config

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type UserInfo struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Token     string `json:"token"`
	ExpiresIn uint64 `json:"expires_in"`
}

type Config struct {
	Idx   int         `json:"user_idx"`
	Users []*UserInfo `json:"users"`
}

// save config
func (c *Config) Save(fname string) error {
	if len(c.Users) == 0 {
		return nil
	}
	fd, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fd.Close()

	var b []byte
	if b, err = json.Marshal(c); err == nil {
		s := base64.StdEncoding.EncodeToString(b)
		_, err = fd.WriteString(s)
	}

	return err
}

// load config
func (c *Config) Load(fname string) error {
	fd, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer fd.Close()

	var b []byte
	if b, err = ioutil.ReadAll(fd); err == nil {
		if b, err = base64.StdEncoding.DecodeString(string(b)); err == nil {
			err = json.Unmarshal(b, c)
		}
	}

	return err
}

// get current user from config
func (c *Config) GetCurUser() *UserInfo {
	if c.Idx >= 0 && c.Idx < len(c.Users) {
		return c.Users[c.Idx]
	}
	return nil
}

// update user info in config
func (c *Config) UpdateUserInfo(u *UserInfo) {
	c.Idx = -1
	for k, v := range c.Users {
		if v.AppKey == u.AppKey {
			c.Idx = k
			break
		}
	}
	if c.Idx == -1 {
		c.Idx = len(c.Users)
		c.Users = append(c.Users, u)
	} else {
		c.Users[c.Idx] = u
	}
}

// switch user by user's appkey from config
func (c *Config) SwitchUser(appKey string) error {
	for k, v := range c.Users {
		if v.AppKey == appKey {
			c.Idx = k
			return nil
		}
	}
	return errors.New("no such user")
}

// remove current user from config
func (c *Config) RemoveUser() error {
	if c.Idx >= 0 && c.Idx < len(c.Users) {
		c.Users = append(c.Users[0:c.Idx], c.Users[c.Idx+1:]...)
		c.Idx = 0
		return nil
	}
	return errors.New("no such user")
}
