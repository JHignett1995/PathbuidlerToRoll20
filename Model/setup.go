package Model

import (
	"encoding/json"
	"io/ioutil"
)

func Setup() (*Config, error) {
	configBytes, err := ioutil.ReadFile("config.json")
	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
