package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DB          DB     `yaml:"db"`
	GitHub      GitHub `yaml:"github"`
	GCPFunction bool   `yaml:"gcp_function"`
}

type FirebaseCredentials struct {
	ProjectID     string `yaml:"project_id"`
	PrivateKeyID  string `yaml:"private_key_id"`
	PrivateKey    string `yaml:"private_key"`
	ClientEmail   string `yaml:"client_email"`
	ClientID      string `yaml:"client_id"`
	ClientCertURL string `yaml:"client_x509_cert_url"`
}

type DB struct {
	Endpoint            string              `yaml:"endpoint"`
	Provider            string              `yaml:"provider"`
	FirebaseCredentials FirebaseCredentials `yaml:"firebase_credentials,omitempty"`
}

type GitHub struct {
	Organization string `yaml:"organization"`
}

// Global configuration for the application.
var configuration Config
var rwMutex sync.RWMutex

func Init(config string) (err error) {
	file := config
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	if file == "" {
		file = pwd + "/configuration.yaml"
	}
	c, err := LoadFromFile(file)
	if err != nil {
		return
	}
	if os.Getenv("GCP_FUNCTION") == "true" {
		c.GCPFunction = true
	}
	Set(c)
	return
}

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
