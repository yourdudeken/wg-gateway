package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/web"
)

var webPort int

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web UI dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		server := web.NewServer("config.yaml")
		
		fmt.Printf("Starting W-G Gateway Web UI...\n")
		fmt.Printf("Dashboard: http://localhost:%d\n", webPort)
		fmt.Printf("Press Ctrl+C to stop\n\n")
		
		if err := server.Start(webPort); err != nil {
			fmt.Printf("Error starting web server: %v\n", err)
		}
	},
}

func init() {
	webCmd.Flags().IntVarP(&webPort, "port", "p", 8080, "Port to run the web UI on")
	rootCmd.AddCommand(webCmd)
}
