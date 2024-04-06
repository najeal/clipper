package specs

import (
	"gopkg.in/yaml.v3"
)

type Root struct {
	Name      string    `yaml:"name"`
	Usage     string    `yaml:"usage"`
	UsageText string    `yaml:"usageText"`
	Action    string    `yaml:"action"`
	Commands  []Command `yaml:"commands"`
}

type Command struct {
	Name        string    `yaml:"name"`
	Usage       string    `yaml:"usage"`
	Action      string    `yaml:"action"`
	SubCommands []Command `yaml:"commands"`
}

func Unmarshal(in []byte) (Root, error) {
	return unmarshalRoot(in)
}

func unmarshalRoot(in []byte) (out Root, err error) {
	err = yaml.Unmarshal(in, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
