package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

const (
	configPath = "config/config.toml"
)

//Config ...
type Config struct {
	DataBase []SQLDataBase `toml:"DataBase"`
}

//SQLDataBase ...
type SQLDataBase struct {
	Server   string `toml:"Server"`
	DataBase string `toml:"DataBase"`
	Port     int    `toml:"Port"`
	User     string
	Password string
}

//Load ...
func (c *Config) Load() error {
	err := c.loadFromFile(configPath)
	if err != nil {
		err = fmt.Errorf("Can't load config from file %v: %w", configPath, err)
		return err
	}

	return nil
}

func (c *Config) loadFromFile(configPath string) error {
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		err = fmt.Errorf("Can't decode config %v: %w", configPath, err)
		return err
	}
	return nil
}
