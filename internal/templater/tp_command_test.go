package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/command/exp_command_template_lvl_1.txt
var expCommandTemplateLvl1 string

//go:embed test_files/command/exp_command_with_flags.txt
var expCommandWithFlagsTemplate string

//go:embed test_files/command/exp_command_template_lvl_2.txt
var expCommandTemplateLvl2 string

//go:embed test_files/command/exp_command_template_lvl_3.txt
var expCommandTemplateLvl3 string

//go:embed test_files/command/exp_command_template_lvl_4.txt
var expCommandTemplateLvl4 string

//go:embed test_files/command/exp_command_template_lvl_1_with_action.txt
var expCommandTemplateLvl1WithAction string

func TestGenerateCommandTemplate(t *testing.T) {
	t.Run("CommandOneLevel", func(t *testing.T) {
		command := Command{
			Name:  "first",
			Usage: "first usage",
		}
		content, err := generateContent("CommandTemplate", commandTemplate, command)
		require.NoError(t, err)
		require.Equal(t, expCommandTemplateLvl1, string(content))

		t.Run("WithAction", func(t *testing.T) {
			command := Command{
				Name:   "first",
				Usage:  "first usage",
				Action: "First",
			}
			content, err := generateContent("CommandTemplate", commandTemplate, command,
				additionalTemplate{"StringFlagTemplate", stringFlagTemplate},
				additionalTemplate{"BoolFlagTemplate", boolFlagTemplate},
			)
			require.NoError(t, err)
			require.Equal(t, expCommandTemplateLvl1WithAction, string(content))
		})
	})
	t.Run("CommandWithFlags", func(t *testing.T) {
		command := Command{
			Name:  "first",
			Usage: "first usage",
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
				{
					Name:  "ages",
					Usage: "ages info",
					Type:  "int64slice",
				},
			},
		}
		content, err := generateContent("CommandTemplate", commandTemplate, command,
			additionalTemplate{"FlagTemplate", flagTemplate},
			additionalTemplate{"StringFlagTemplate", stringFlagTemplate},
			additionalTemplate{"BoolFlagTemplate", boolFlagTemplate},
			additionalTemplate{"Int64FlagTemplate", int64FlagTemplate},
			additionalTemplate{"Int64SliceFlagTemplate", int64sliceFlagTemplate},
		)
		require.NoError(t, err)
		require.Equal(t, expCommandWithFlagsTemplate, string(content))
	})
	t.Run("CommandTwoLevel", func(t *testing.T) {
		command := Command{
			Name:  "first",
			Usage: "first usage",
			SubCommands: []Command{
				{
					Name:  "second",
					Usage: "second usage",
				},
			},
		}
		content, err := generateContent("CommandTemplate", commandTemplate, command)
		require.NoError(t, err)
		require.Equal(t, expCommandTemplateLvl2, string(content))
	})
	t.Run("CommandThreeLevel", func(t *testing.T) {
		command := Command{
			Name:  "first",
			Usage: "first usage",
			SubCommands: []Command{
				{
					Name:  "second",
					Usage: "second usage",
					SubCommands: []Command{
						{
							Name:  "third",
							Usage: "third usage",
						},
					},
				},
			},
		}
		content, err := generateContent("CommandTemplate", commandTemplate, command)
		require.NoError(t, err)
		require.Equal(t, expCommandTemplateLvl3, string(content))
	})
	t.Run("CommandFourLevel", func(t *testing.T) {
		command := Command{
			Name:  "first",
			Usage: "first usage",
			SubCommands: []Command{
				{
					Name:  "second",
					Usage: "second usage",
					SubCommands: []Command{
						{
							Name:  "third",
							Usage: "third usage",
							SubCommands: []Command{
								{
									Name:  "fourth",
									Usage: "fourth usage",
								},
							},
						},
					},
				},
			},
		}
		content, err := generateContent("CommandTemplate", commandTemplate, command)
		require.NoError(t, err)
		require.Equal(t, expCommandTemplateLvl4, string(content))
	})
}
