package specs

import (
	"gopkg.in/yaml.v3"
)

type Root struct {
	Name        string    `yaml:"name"`
	Version     string    `yaml:"version"`
	VersionFlag Flag      `yaml:"versionFlag"`
	Usage       string    `yaml:"usage"`
	UsageText   string    `yaml:"usageText"`
	Action      string    `yaml:"action"`
	Flags       []Flag    `yaml:"flags"`
	Commands    []Command `yaml:"commands"`
}

type Command struct {
	Name        string    `yaml:"name"`
	Usage       string    `yaml:"usage"`
	Action      string    `yaml:"action"`
	Flags       []Flag    `yaml:"flags"`
	SubCommands []Command `yaml:"commands"`
}

type Flag struct {
	Name    string   `yaml:"name"`
	Value   string   `yaml:"value"`
	Type    string   `yaml:"type"`
	Usage   string   `yaml:"usage"`
	Aliases []string `yaml:"aliases"`
	EnvVars []string `yaml:"envVars"`
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
