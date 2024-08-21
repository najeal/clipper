package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/cmd/clipper/gen"
	"github.com/najeal/clipper/icobra"
	"github.com/urfave/cli/v2"
)

type Service struct{}

func (*Service) Run(ctx *cli.Context) error {
	const packageName = "cmd"
	specPath := ctx.Args().Get(0)

	specFile, err := os.ReadFile(specPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read spec file: %v", err)
		return gen.ExitReadManifestFailed
	}

	//	var generator cliGenerator = pkg.NewCliGenerator()
	//	codeFile, err := generator.Generate(specFile, packageName)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "failed to generate cli code file: %v\n", err)
	//		return gen.ExitCliGenerationFailed
	//	}

	//	fmt.Fprintln(os.Stdout, string(codeFile))
	//	return nil
	cligen := icobra.NewCliGenerator()
	_, err = cligen.Generate(specFile, packageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate cli code file: %v\n", err)
		return gen.ExitCliGenerationFailed
	}
	return nil
}

func main() {
	app := gen.NewApp(&Service{})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
