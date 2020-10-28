package config

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

const (
	configPath = "config/config.toml"
)

//Config ...
type Config struct {
	DataBase  []SQLDataBase `toml:"DataBase"`
	ServerOpt ServerOpt     `toml:"ServerOpt"`
}

//SQLDataBase ...
type SQLDataBase struct {
	Server     string `toml:"Server"`
	DataBase   string `toml:"DataBase"`
	Port       int    `toml:"Port"`
	User       string `toml:"User"`
	Password   string `toml:"Password"`
	SearchPath string `toml:"SearchPath"`
	Driver     string `toml:"Driver"`
}

//ServerOpt ...
type ServerOpt struct {
	Port           string   `toml:"Port"`
	AllowedHeaders []string `toms:"AllowedHeaders"`
	ExposedHeaders []string `toms:"ExposedHeaders"`
	AllowedMethods []string `toms:"AllowedMethods"`
}

type duration time.Duration

//Load ....
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
