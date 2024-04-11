package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/root/exp_root.txt
var expRoot string

//go:embed test_files/root/exp_root_complete.txt
var expRootComplete string

//go:embed test_files/root/exp_root_with_version.txt
var expRootWithVersion string

//go:embed test_files/root/exp_root_with_exits.txt
var expRootWithExits string

func TestGenerateRootTemplate(t *testing.T) {
	t.Run("Root", func(t *testing.T) {
		app := Root{
			PackageName: "genroot",
			App:         App{},
		}
		content, err := generateContent("RootTemplate", rootTemplate, app,
			additionalTemplate{"NewAppTemplate", newAppTemplate}, additionalTemplate{"CommandTemplate", commandTemplate})
		require.NoError(t, err)
		require.Equal(t, expRoot, string(content))
	})
	t.Run("WithExits", func(t *testing.T) {
		app := Root{
			PackageName: "genroot",
			App:         App{},
			Exits: []Exit{
				{
					VarName: "ExitUserNotFound",
					Message: "user has not been found",
					Code:    1,
				},
				{
					VarName: "ExitTVNotFound",
					Message: "tv has not been found",
					Code:    2,
				},
			},
		}
		content, err := generateContent("RootTemplate", rootTemplate, app,
			additionalTemplate{"NewAppTemplate", newAppTemplate},
			additionalTemplate{"CommandTemplate", commandTemplate},
			additionalTemplate{"ExitTemplate", exitTemplate},
		)
		require.NoError(t, err)
		require.Equal(t, expRootWithExits, string(content))
	})
	t.Run("RootWithVersion", func(t *testing.T) {
		app := Root{
			PackageName: "genroot",
			Version:     "v0.1.0",
			VersionFlag: Flag{
				Name:    "version",
				Usage:   "print the version",
				Aliases: []string{"v"},
			},
			App: App{},
		}
		content, err := generateContent("RootTemplate", rootTemplate, app,
			additionalTemplate{"NewAppTemplate", newAppTemplate},
			additionalTemplate{"CommandTemplate", commandTemplate},
			additionalTemplate{"BoolFlagTemplate", boolFlagTemplate},
		)
		require.NoError(t, err)
		require.Equal(t, expRootWithVersion, string(content))
	})
	t.Run("Complete", func(t *testing.T) {
		app := Root{
			PackageName: "genroot",
			Methods:     []string{"Move"},
			App: App{
				Name:  "app",
				Usage: "app usage",
				Commands: []Command{
					{
						Name:   "move",
						Usage:  "move usage",
						Action: "Move",
					},
				},
			},
		}
		content, err := generateContent("RootTemplate", rootTemplate, app,
			additionalTemplate{"NewAppTemplate", newAppTemplate}, additionalTemplate{"CommandTemplate", commandTemplate})
		require.NoError(t, err)
		require.Equal(t, expRootComplete, string(content))
	})
}
