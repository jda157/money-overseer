package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config struct {
	TelegramToken string `yaml:"telegram_token"`
}

func InitConfig(p string) (*config, error) {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("can't read from file configutration file:%v", err)
	}
	a := &config{}
	err = yaml.Unmarshal(file, &a)
	return a, nil
}

func (c *config) GetTelegramToken() string {
	return c.TelegramToken
}
