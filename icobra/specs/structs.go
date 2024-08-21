package specs

import "gopkg.in/yaml.v3"

type RootApp struct {
	Name     string    `yaml:"name"`
	Use      string    `yaml:"use"`
	Short    string    `yaml:"short"`
	Long     string    `yaml:"long"`
	Runnable bool      `yaml:"runnable"`
	Flags    []Flag    `yaml:"flags"`
	Commands []Command `yaml:"commands"`
}

type Command struct {
	CommandName string    `yaml:"name"`
	Use         string    `yaml:"use"`
	Short       string    `yaml:"short"`
	Long        string    `yaml:"long"`
	Runnable    bool      `yaml:"runnable"`
	Flags       []Flag    `yaml:"flags"`
	SubCommands []Command `yaml:"commands"`
}

type Flag struct {
	Name       string
	Type       string
	Usage      string
	Value      string
	Shorthand  string
	Persistent bool
}

func Unmarshal(in []byte) (RootApp, error) {
	return unmarshalRootApp(in)
}

func unmarshalRootApp(in []byte) (out RootApp, err error) {
	err = yaml.Unmarshal(in, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
