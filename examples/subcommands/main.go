package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/examples/subcommands/gen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := gen.NewApp(&Service{})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Add(_ *cli.Context) error {
	fmt.Fprintln(os.Stdout, "Add command !")
	return nil
}

func (*Service) Complete(_ *cli.Context) error {
	fmt.Fprintln(os.Stdout, "Complete command !")
	return nil
}

func (*Service) TemplateAdd(_ *cli.Context) error {
	fmt.Fprintln(os.Stdout, "TemplateAdd command !")
	return nil
}

func (*Service) TemplateRemove(_ *cli.Context) error {
	fmt.Fprintln(os.Stdout, "TemplateRemove command !")
	return nil
}
