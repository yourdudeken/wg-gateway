package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Clean up generated deployment files",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.RemoveAll("deploy")
		if err != nil {
			fmt.Printf("Error removing deploy directory: %v\n", err)
			return
		}

		fmt.Println("Cleaned up 'deploy' directory.")
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
