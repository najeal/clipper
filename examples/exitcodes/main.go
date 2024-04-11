package main

import (
	"log"
	"os"

	"github.com/najeal/clipper/examples/exitcodes/gen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := gen.NewApp(&Service{})
	app.RunAndExitOnError()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Run(cCtx *cli.Context) error {
	if !cCtx.Bool("ginger-crouton") {
		return cli.Exit("Ginger croutons are not in the soup", 86)
	}
	return nil
}
