package main

import (
	"fmt"
	"log"
	"os"

	pkg "github.com/najeal/clipper/internal"
	"github.com/urfave/cli/v2"
)

const (
	version = "0.1.0"
)

type cliGenerator interface {
	Generate(specFile []byte, packageName string) ([]byte, error)
}

func init() {
}

func main() {
	app := &cli.App{
		Version:   version,
		Name:      "clipper",
		Usage:     "generate urfave/cli code from yaml specifications",
		UsageText: "clipper SPECS_SRC_PATH > DEST_GO_FILE_PATH",
		Action: func(cliCtx *cli.Context) error {
			const packageName = "gen"
			specPath := cliCtx.Args().Get(0)

			specFile, err := os.ReadFile(specPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read spec file: %v", err)
				os.Exit(1)
			}

			var generator cliGenerator = pkg.NewCliGenerator()
			codeFile, err := generator.Generate(specFile, packageName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to generate cli code file: %v\n", err)
				os.Exit(1)
			}

			fmt.Fprintln(os.Stdout, string(codeFile))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
