package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/examples/flags/gen"
	"github.com/urfave/cli/v2"
)

func main() {
	app := gen.NewApp(&Service{})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Run(cCtx *cli.Context) error {
	name := "Nefertiti"
	if cCtx.NArg() > 0 {
		name = cCtx.Args().Get(0)
	}
	if cCtx.String("lang") == "spanish" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println("Hello", name)
	}
	return nil
}
