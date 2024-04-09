// Code generated by clipper. DO NOT EDIT.
package gen

import (
	"github.com/urfave/cli/v2"
)

type Service interface {
	Add(*cli.Context) error
	Complete(*cli.Context) error
	TemplateAdd(*cli.Context) error
	TemplateRemove(*cli.Context) error
}

func NewApp(svc Service) *cli.App {
	var app = &cli.App{
		Name:  "subcommand-manager",
		Usage: "",
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a task to the list",
				Action: func(cliCtx *cli.Context) error {
					return svc.Add(cliCtx)
				},
			},
			{
				Name:  "complete",
				Usage: "complete a task on the list",
				Action: func(cliCtx *cli.Context) error {
					return svc.Complete(cliCtx)
				},
			},
			{
				Name:  "template",
				Usage: "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cliCtx *cli.Context) error {
							return svc.TemplateAdd(cliCtx)
						},
					}, {
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cliCtx *cli.Context) error {
							return svc.TemplateRemove(cliCtx)
						},
					},
				},
			},
		},
	}
	return app
}
