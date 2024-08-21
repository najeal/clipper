// Code generated by clipper. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

var templateAddCmd *cobra.Command

func NewTemplateAddCmd(svc Service) *cobra.Command {
	templateAddCmd = &cobra.Command{
		Use:   "add",
		Short: "add a new template",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			svc.TemplateAdd(cmd, args)
		},
	}
	return templateAddCmd
}
