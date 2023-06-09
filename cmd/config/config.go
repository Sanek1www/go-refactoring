package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	AppPort         string `yaml:"port"`
	JsonStorageLink string `yaml:"json_storage_link"`
}

func Parse(confPath string) (*Config, error) {
	filename, err := filepath.Abs(confPath)
	if err != nil {
		return nil, fmt.Errorf("can't get config path: %s", err.Error())
	}

	yamlConf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read conf: %s", err.Error())
	}

	var config Config
	err = yaml.Unmarshal(yamlConf, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshall conf: %s", err.Error())
	}

	return &config, nil
}
