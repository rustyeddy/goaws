package cmd

import "github.com/spf13/cobra"

var (
	regionCmd = cobra.Command{
		Use:   "region list",
		Short: "list regions a select which regions to work with",
		Long:  "AWS divides most things into regions, manage them",
	}
)

func init() {

}
