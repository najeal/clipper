package main

import (
	"fmt"
	"os"

	"github.com/najeal/clipper/examples/cobra_flags/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute(&Service{})
}

type Service struct{}

func (*Service) Add(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Add command!")
}

func (*Service) Complete(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "Complete command !")
}

func (*Service) TemplateAdd(cmd *cobra.Command, args []string) {
	lang, err := cmd.Flags().GetString("lang")
	if err != nil {
		panic(err.Error())
	}
	number, err := cmd.Flags().GetInt64("number")
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "TemplateAdd command %d %s!\n", number, lang)
	ports, err := cmd.Flags().GetInt64Slice("ports")
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "possible ports:\n")
	for _, port := range ports {
		fmt.Fprintf(os.Stdout, "%d\n", port)
	}
}

func (*Service) TemplateRemove(cmd *cobra.Command, args []string) {
	fmt.Fprintln(os.Stdout, "TemplateRemove command !")
}
