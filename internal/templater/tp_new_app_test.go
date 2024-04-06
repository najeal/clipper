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
