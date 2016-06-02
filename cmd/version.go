package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of the Kubernetes Cloud Foundry TPS.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tpsVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
