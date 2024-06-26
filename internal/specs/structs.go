package specs

import (
	"gopkg.in/yaml.v3"
)

type Root struct {
	Name                   string              `yaml:"name"`
	Version                string              `yaml:"version"`
	VersionFlag            Flag                `yaml:"versionFlag"`
	Copyright              string              `yaml:"copyright"`
	Usage                  string              `yaml:"usage"`
	UsageText              string              `yaml:"usageText"`
	Action                 string              `yaml:"action"`
	Flags                  []Flag              `yaml:"flags"`
	UseShortOptionHandling bool                `yaml:"useShortOptionHandling"`
	EnableBashCompletion   bool                `yaml:"enableBashCompletion"`
	Commands               []Command           `yaml:"commands"`
	ExitCodes              map[string]ExitCode `yaml:"exitCodes"`
}

type ExitCode struct {
	Message string `yaml:"message"`
	Code    int    `yaml:"code"`
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
