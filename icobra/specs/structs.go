package specs

import "gopkg.in/yaml.v3"

type RootApp struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	Usage    string `yaml:"usage"`
	Action   string `yaml:"action"`
	Short    string `yaml:"short"`
	Long     string `yaml:"long"`
	Runnable bool   `yaml:"runnable"`
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
