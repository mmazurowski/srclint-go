package config

import (
	"os"

	"srclint/src/errors"

	"gopkg.in/yaml.v3"
)

type RawConfig struct {
	Version   int      `yaml:"version"`
	Structure []string `yaml:"structure"`
	Required  []string `yaml:"required"`
	Ignore    []string `yaml:"ignore"`
}

func ReadConfig(path string) (config *RawConfig, configError error) {
	var data, err = os.ReadFile(path)

	if err != nil {
		return nil, errors.CreateConfigParserError("Could not read config file.")
	}

	var configContent RawConfig
	err = yaml.Unmarshal(data, &configContent)

	if err != nil {
		return nil, errors.CreateConfigParserError("Provided config is invalid")
	}

	return &configContent, nil
}
