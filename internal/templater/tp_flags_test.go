package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/flags/exp_bool1.txt
var expBoolFlag1 string

//go:embed test_files/flags/exp_bool_no_aliases.txt
var expBoolFlagNoAliases string

//go:embed test_files/flags/exp_bool_no_envvars.txt
var expBoolFlagNoEnvVars string

//go:embed test_files/flags/exp_bool_no_value.txt
var expBoolFlagNoValue string

//go:embed test_files/flags/exp_string1.txt
var expStringFlag1 string

//go:embed test_files/flags/exp_string_no_aliases.txt
var expStringFlagNoAliases string

//go:embed test_files/flags/exp_string_no_envvars.txt
var expStringFlagNoEnvVars string

//go:embed test_files/flags/exp_string_no_value.txt
var expStringFlagNoValue string

//go:embed test_files/flags/exp_int64.txt
var expInt64Flag1 string

func TestGenerateFlagTemplate(t *testing.T) {
	t.Run("Int64", func(t *testing.T) {
		t.Run("Complete", func(t *testing.T) {
			flag := Flag{
				Name:    "size",
				Value:   "4",
				Type:    "int64",
				Usage:   "size info",
				Aliases: []string{"s"},
				EnvVars: []string{"F_SIZE"},
			}

			content, err := generateContent("Int64FlagTemplate", int64FlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expInt64Flag1, string(content))
		})
	})
	t.Run("Bool", func(t *testing.T) {
		t.Run("Complete", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "true",
				Type:    "bool",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("BoolFlagTemplate", boolFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expBoolFlag1, string(content))
		})
		t.Run("NoAliases", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "true",
				Type:    "bool",
				Usage:   "forcing stuff",
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("BoolFlagTemplate", boolFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expBoolFlagNoAliases, string(content))
		})
		t.Run("NoEnvVars", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "true",
				Type:    "bool",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
			}

			content, err := generateContent("BoolFlagTemplate", boolFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expBoolFlagNoEnvVars, string(content))
		})
		t.Run("NoValue", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Type:    "bool",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("BoolFlagTemplate", boolFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expBoolFlagNoValue, string(content))
		})
	})
	t.Run("String", func(t *testing.T) {
		t.Run("Complete", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "forcing",
				Type:    "string",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("StringFlagTemplate", stringFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expStringFlag1, string(content))
		})
		t.Run("NoAliases", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "forcing",
				Type:    "bool",
				Usage:   "forcing stuff",
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("StringFlagTemplate", stringFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expStringFlagNoAliases, string(content))
		})
		t.Run("NoEnvVars", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Value:   "forcing",
				Type:    "bool",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
			}

			content, err := generateContent("StringFlagTemplate", stringFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expStringFlagNoEnvVars, string(content))
		})
		t.Run("NoValue", func(t *testing.T) {
			flag := Flag{
				Name:    "force",
				Type:    "bool",
				Usage:   "forcing stuff",
				Aliases: []string{"f"},
				EnvVars: []string{"CP_FORCE"},
			}

			content, err := generateContent("BoolFlagTemplate", stringFlagTemplate, flag)
			require.NoError(t, err)
			require.Equal(t, expStringFlagNoValue, string(content))
		})
	})
}
