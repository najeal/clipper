// Code generated by clipper. DO NOT EDIT.
package gen

import (
	"github.com/urfave/cli/v2"
)

type Service interface {
	Run(*cli.Context) error
}

var (
	ExitReadManifestFailed  = cli.Exit("failed to read manifest", 1)
	ExitCliGenerationFailed = cli.Exit("failed to generate cli", 2)
)

var (
	Version = "0.1.0"
)

func init() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print version",
	}
}

func NewApp(svc Service) *cli.App {
	var app = &cli.App{
		Name:    "clipper",
		Version: Version,
		Usage:   "clipper generates cli from manifest",
		Action: func(cliCtx *cli.Context) error {
			return svc.Run(cliCtx)
		},
	}
	return app
}
