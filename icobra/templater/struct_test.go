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

func TestGenerateCommandFile(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		cmdData := Command{
			CommandName: "ConfigCreate",
			VarCmd:      "configCreate",
			Use:         "create",
			Short:       "create desc",
			Long:        "create desc",
			Run:         true,
		}

		res, err := GenerateCommand(cmdData)
		require.NoError(t, err)
		require.Equal(t, expectedCommandFileBasic, string(res))
	})
}
