// Code generated by clipper. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

var templateCmd *cobra.Command

func NewTemplateCmd(svc Service) *cobra.Command {
	templateCmd = &cobra.Command{
		Use:   "template",
		Short: "options for task templates",
		Long:  ``,
	}
	return templateCmd
}
