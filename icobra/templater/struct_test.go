package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/exp_main.txt
var expectedMain string

//go:embed test_files/exp_root_basic.txt
var expectedRootBasic string

//go:embed test_files/exp_root_run.txt
var expectedRootRun string

//go:embed test_files/exp_command_file_basic.txt
var expectedCommandFileBasic string

//go:embed test_files/exp_command_file_with_flags.txt
var expectedCommandFileWithFlags string

func TestGenerateMain(t *testing.T) {
	mainData := MainData{
		ProjectRoot: "/my-project",
	}
	res, err := GenerateMain(mainData)
	require.NoError(t, err)
	require.Equal(t, expectedMain, string(res))
}

func TestGenerateRootFile(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		rootData := RootData{
			Command: Command{
				CommandName: "App",
				VarCmd:      "app",

				Use:   "app",
				Short: "app short description",
				Long:  "app long description",
				Run:   false,
			},
			Methods: []string{
				"App", "Config", "ConfigCreate", "Fetch",
			},
			CommandsHierarchy: []Hierarchy{
				{
					Parent: "app",
					Child:  "Config",
				},
				{
					Parent: "app",
					Child:  "Fetch",
				},
				{
					Parent: "config",
					Child:  "ConfigCreate",
				},
			},
		}
		res, err := GenerateRoot(rootData)
		require.NoError(t, err)
		require.Equal(t, expectedRootBasic, string(res))
	})
	t.Run("Run", func(t *testing.T) {
		rootData := RootData{
			Command: Command{
				CommandName: "App",
				VarCmd:      "app",

				Use:   "app",
				Short: "app short description",
				Long:  "app long description",
				Run:   true,
			},
		}
		res, err := GenerateRoot(rootData)
		require.NoError(t, err)
		require.Equal(t, expectedRootRun, string(res))
	})
}

func TestGenerateFlags(t *testing.T) {
	t.Run("Int64", func(t *testing.T) {
		flag := Flag{
			Name:  "age",
			Type:  "int64",
			Value: "50",
			Usage: "indicate people age",
		}
		expectedRes := []byte(`Flags().Int64P("age", "", 50, "indicate people age")`)
		res, err := generateFlagContent(flag)
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})
}

func TestGenerateCommandFile(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		cmdData := Command{
			CommandName: "ConfigCreate",
			VarCmd:      "configCreate",
			Use:         "create",
			Short:       "create desc",
			Long:        "create desc",
			Run:         true,
			Flags: []Flag{
				{
					Name:       "max",
					Usage:      "indicate max",
					Type:       "int64",
					Shorthand:  "m",
					Value:      "50",
					Persistent: true,
					VarCmd:     "configCreate",
				},
			},
		}

		res, err := GenerateCommand(cmdData)
		require.NoError(t, err)
		require.Equal(t, expectedCommandFileBasic, string(res))
	})

	t.Run("WithFlags", func(t *testing.T) {
		cmdData := Command{
			CommandName: "ConfigCreate",
			VarCmd:      "configCreate",
			Use:         "create",
			Short:       "create desc",
			Long:        "create desc",
			Run:         true,
			Flags: []Flag{
				{
					Name:       "name",
					Usage:      "naming",
					Type:       "string",
					Shorthand:  "n",
					Value:      "myname",
					Persistent: false,
					VarCmd:     "configCreate",
				},
				{
					Name:       "ports",
					Usage:      "possible ports",
					Type:       "int64Slice",
					Shorthand:  "p",
					Value:      "80,8080,3000",
					Persistent: false,
					VarCmd:     "configCreate",
				},
				{
					Name:       "validation",
					Usage:      "enable validation",
					Type:       "bool",
					Shorthand:  "v",
					Value:      "true",
					Persistent: false,
					VarCmd:     "configCreate",
				},
				{
					Name:       "min",
					Usage:      "indicate min",
					Type:       "int",
					Shorthand:  "",
					Value:      "10",
					Persistent: false,
					VarCmd:     "configCreate",
				},
				{
					Name:       "max",
					Usage:      "indicate max",
					Type:       "int64",
					Shorthand:  "m",
					Value:      "50",
					Persistent: true,
					VarCmd:     "configCreate",
				},
			},
		}

		res, err := GenerateCommand(cmdData)
		require.NoError(t, err)
		require.Equal(t, expectedCommandFileWithFlags, string(res))
	})
}
