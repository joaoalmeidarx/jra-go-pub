package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Appname  string `json:"appname"`
	Version  string `json:"version"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

var Default *Config

func LoadDefault(filename string) error {
	configFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	byteConfigFile, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}
	configFile.Close()

	var config = &Config{}
	err = json.Unmarshal(byteConfigFile, &config)
	if err != nil {
		return err
	}

	Default = config
	return nil
}
