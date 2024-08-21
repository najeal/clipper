package main

import (
	"fmt"
	"os"

	"github.com/najeal/clipper/examples/cobra_subcommands/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute(&Service{})
}

type Service struct{}

func (*Service) Add(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Add command !")
}

func (*Service) Complete(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Complete command !")
}

func (*Service) TemplateAdd(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "TemplateAdd command !")
}

func (*Service) TemplateRemove(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "TemplateRemove command !")
}
