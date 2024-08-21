package icobra

import (
	"strings"

	"github.com/najeal/clipper/icobra/specs"
	"github.com/najeal/clipper/icobra/templater"
)

func transform(packageName string, in specs.RootApp) templater.RootData {
	return templater.RootData{
		Methods: []string{in.Name},
		Command: templater.Command{
			Use:         in.Usage,
			Short:       in.Short,
			Long:        in.Long,
			CommandName: in.Name,
			VarCmd:      strings.ToLower(string(in.Name[0])) + in.Name[1:],
			Run:         in.Runnable,
		},
	}
}
