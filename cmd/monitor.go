package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/monitor"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start the proactive gateway monitor",
	Long: `Starts the background monitoring service that periodically checks:
- VPS connectivity
- Peer availability
- Service port responsiveness
And sends alerts via Discord or Telegram if anomalies are detected.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		if !cfg.Monitor.Discord.Enabled && !cfg.Monitor.Telegram.Enabled {
			fmt.Println("Warning: No alert channels (Discord/Telegram) are enabled in config.yaml.")
			fmt.Println("The monitor will log to console only.")
		}

		hub := monitor.NewHub(cfg)
		fmt.Println("W-G Gateway Monitor Service")
		fmt.Println("---------------------------")
		hub.Start()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
