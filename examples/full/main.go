package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/examples/full/gen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := gen.NewApp(&Service{})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Greet(ctx *cli.Context) error {
	fmt.Fprintln(os.Stdout, "greeting !")
	return nil
}

func (*Service) Check(ctx *cli.Context) error {
	if !ctx.Bool("force") {
		return gen.ExitMissingForce
	}
	fmt.Fprintln(os.Stdout, "force flag detected")
	return nil
}
