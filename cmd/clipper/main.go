package main

import (
	"fmt"
	"log"
	"os"

	"github.com/najeal/clipper/cmd/clipper/gen"
	pkg "github.com/najeal/clipper/internal"
	"github.com/urfave/cli/v2"
)

type cliGenerator interface {
	Generate(specFile []byte, packageName string) ([]byte, error)
}

type Service struct{}

func (*Service) Run(ctx *cli.Context) error {
	const packageName = "gen"
	specPath := ctx.Args().Get(0)

	specFile, err := os.ReadFile(specPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read spec file: %v", err)
		return gen.ExitReadManifestFailed
	}

	var generator cliGenerator = pkg.NewCliGenerator()
	codeFile, err := generator.Generate(specFile, packageName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate cli code file: %v\n", err)
		return gen.ExitCliGenerationFailed
	}

	fmt.Fprintln(os.Stdout, string(codeFile))
	return nil
}

func main() {
	app := gen.NewApp(&Service{})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
