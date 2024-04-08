package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/examples/greet/gen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := gen.NewApp(&Service{})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Greet(_ *cli.Context) error {
	fmt.Fprintln(os.Stdout, "Hello friend!")
	return nil
}
