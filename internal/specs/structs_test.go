package specs

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

//go:embed test_files/cmd1.yaml
var cmd1 string

//go:embed test_files/cmd_flags.yaml
var cmdFlags string

func TestUnmarshal(t *testing.T) {
	t.Run("Cmd1", func(t *testing.T) {
		root, err := unmarshalRoot([]byte(cmd1))
		require.NoError(t, err)
		require.Equal(t, Root{
			Name:      "cp",
			Usage:     "cp is copying a file",
			UsageText: "cp SRC_PATH DEST_PATH",
			Action:    "Copy",
		}, root)
	})
	t.Run("CmdFlags", func(t *testing.T) {
		root, err := unmarshalRoot([]byte(cmdFlags))
		require.NoError(t, err)
		require.Equal(t, Root{
			Name:      "cp",
			Usage:     "cp is copying a file",
			UsageText: "cp SRC_PATH DEST_PATH",
			Action:    "Copy",
			Flags: []Flag{
				{
					Name:    "force",
					Type:    "bool",
					Value:   "true",
					Aliases: []string{"f"},
					Usage:   "If the destination file cannot be opened, remove it and create a new file",
					EnvVars: []string{"COPY_FORCE"},
				},
				{
					Name:  "verbose",
					Type:  "string",
					Value: "info",
					Usage: "choose verbosity",
				},
				{
					Name:  "age",
					Type:  "int64",
					Value: "",
				},
			},
		}, root)
	})
}

func TestFlagsUnmarshal(t *testing.T) {
	flag := Flag{}
	in := []byte(`
name: toto
type: string
usage: my usage
aliases:
  - t
envVars:
  - TOTO_VAR`)

	err := yaml.Unmarshal(in, &flag)
	require.NoError(t, err)
	require.Equal(t, "toto", flag.Name)
	require.Equal(t, "string", flag.Type)
	require.Equal(t, "my usage", flag.Usage)
	require.Equal(t, []string{"t"}, flag.Aliases)
	require.Equal(t, []string{"TOTO_VAR"}, flag.EnvVars)
}
