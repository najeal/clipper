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
