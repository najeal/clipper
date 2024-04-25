package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/new_app/exp_new_app.txt
var expNewApp string

//go:embed test_files/new_app/exp_new_app_with_action.txt
var expNewAppWithAction string

//go:embed test_files/new_app/exp_new_app_with_commands.txt
var expNewAppWithCommands string

//go:embed test_files/new_app/exp_new_app_with_flags.txt
var expNewAppWithFlags string

//go:embed test_files/new_app/exp_new_app_with_version.txt
var expNewAppWithVersion string

//go:embed test_files/new_app/exp_new_app_with_copyright.txt
var expNewAppWithCopyright string

//go:embed test_files/new_app/exp_new_app_with_short_option_handling.txt
var expNewAppWithShortOptionHandling string

//go:embed test_files/new_app/exp_new_app_with_bash_completion.txt
var expNewAppWithBashCompletion string

func TestGenerateNewAppTemplate(t *testing.T) {
	t.Run("Root", func(t *testing.T) {
		app := App{
			Name:  "app",
			Usage: "app usage",
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewApp, string(content))
	})
	t.Run("WithBashCompletion", func(t *testing.T) {
		app := App{
			Name:                 "withBashCompletionHandling",
			Usage:                "using version flag",
			EnableBashCompletion: true,
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithBashCompletion, string(content))
	})
	t.Run("WithShortOptionHandling", func(t *testing.T) {
		app := App{
			Name:                   "withShortOptionHandling",
			Usage:                  "using version flag",
			UseShortOptionHandling: true,
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithShortOptionHandling, string(content))
	})
	t.Run("WithVersion", func(t *testing.T) {
		app := App{
			Name:    "withversion",
			Usage:   "using version flag",
			Version: true,
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithVersion, string(content))
	})
	t.Run("WithCopyright", func(t *testing.T) {
		app := App{
			Name:      "withversion",
			Usage:     "using version flag",
			Copyright: "(c) 2024 Clipper",
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithCopyright, string(content))
	})

	t.Run("WithAction", func(t *testing.T) {
		app := App{
			Name:   "app",
			Usage:  "app usage",
			Action: "App",
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithAction, string(content))
	})
	t.Run("WithFlags", func(t *testing.T) {
		app := App{
			Name:  "app",
			Usage: "app usage",
			Flags: []Flag{
				{
					Name:  "force",
					Usage: "forcing stuff",
					Type:  "bool",
				},
				{
					Name:  "firstname",
					Usage: "gives user first name",
					Type:  "string",
				},
				{
					Name:  "size",
					Usage: "size info",
					Type:  "int64",
				},
			},
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app,
			additionalTemplate{"FlagTemplate", flagTemplate},
			additionalTemplate{"StringFlagTemplate", stringFlagTemplate},
			additionalTemplate{"BoolFlagTemplate", boolFlagTemplate},
			additionalTemplate{"Int64FlagTemplate", int64FlagTemplate},
		)
		require.NoError(t, err)
		require.Equal(t, expNewAppWithFlags, string(content))
	})
	t.Run("WithCommands", func(t *testing.T) {
		app := App{
			Name:  "app",
			Usage: "app usage",
			Commands: []Command{
				{
					Name:   "first",
					Usage:  "first usage",
					Action: "First",
					SubCommands: []Command{
						{
							Name:   "second",
							Usage:  "second usage",
							Action: "Second",
						},
					},
				},
			},
		}
		content, err := generateContent("NewAppTemplate", newAppTemplate, app, additionalTemplate{"CommandTemplate", commandTemplate})
		require.NoError(t, err)
		require.Equal(t, expNewAppWithCommands, string(content))
	})
}
