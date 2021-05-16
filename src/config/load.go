package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// LoadFromFile - загружает конфиг из файла в структуру.
func LoadFromFile(configPath string, config interface{}) (err error) {
	if configPath == "" {
		configPath = defaultPath
	}

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return
	}

	return yaml.Unmarshal(configData, config)
}
