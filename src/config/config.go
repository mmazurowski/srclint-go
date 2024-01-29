package config

import (
	"os"

	"srclint/src/errors"
	"srclint/src/util"

	"gopkg.in/yaml.v3"
)

type RawConfig struct {
	Version   int      `yaml:"version"`
	Structure []string `yaml:"structure"`
	Required  []string `yaml:"required"`
	Ignore    []string `yaml:"ignore"`
}

func CreateEmpty() RawConfig {
	return RawConfig{
		Version:   1,
		Structure: make([]string, 0),
		Required:  make([]string, 0),
		Ignore:    make([]string, 0),
	}
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

// WriteConfig Creates srclint.yml if not exist in given path.
func WriteConfig(cfn RawConfig, path string) {
	info, fileErr := os.Stat(path)

	if fileErr == nil && info.Size() > 0 {
		return
	}

	content, err := yaml.Marshal(cfn)

	util.CheckExit(err, 1)

	err = os.WriteFile(path, content, 0644)

	util.CheckExit(err, 1)
}
