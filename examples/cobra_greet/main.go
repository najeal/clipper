package main

import (
	"fmt"
	"os"

	"github.com/najeal/clipper/examples/cobra_greet/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute(&Service{})
}

type Service struct{}

func (*Service) Greet(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Hello friend!")
}
