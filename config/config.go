package config

import (
	"fmt"
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DB     DB     `yaml:"db"`
	GitHub GitHub `yaml:"github"`
}

type DB struct {
	Endpoint   string `yaml:"endpoint"`
	FB_account string `yaml:"FB_account"`
}

type GitHub struct {
	Organization string `yaml:"organization"`
}

// Global configuration for the application.
var configuration Config
var rwMutex sync.RWMutex

func Get() (conf *Config) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	copy := configuration
	return &copy
}

func Set(conf *Config) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	configuration = *conf
}

func LoadFromFile(filename string) (conf *Config, err error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file [%v]. error=%v", filename, err)
	}

	err = yaml.Unmarshal([]byte(fileContent), &conf)
	if err != nil {
		return
	}
	return
}
