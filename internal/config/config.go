package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type telegram struct {
	TelegramToken string `yaml:"telegram_token"`
}

type database struct {
	Address  string `yaml:"address"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type config struct {
	Telegram telegram
	Database database
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
	return c.Telegram.TelegramToken
}

func (c *config) GetDatabaseLogin() string {
	return c.Database.Login
}

func (c *config) GetDatabasePassword() string {
	return c.Database.Password
}

func (c *config) GetDatabasePort() string {
	return c.Database.Port
}
