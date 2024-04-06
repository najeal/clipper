package specs

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/cmd1.yaml
var commandTemplate string

func TestUnmarshal(t *testing.T) {
	root, err := unmarshalRoot([]byte(commandTemplate))
	require.NoError(t, err)
	require.Equal(t, Root{
		Name:      "cp",
		Usage:     "cp is copying a file",
		UsageText: "cp SRC_PATH DEST_PATH",
		Action:    "Copy",
	}, root)
}
