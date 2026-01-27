package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current configuration status",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		fmt.Printf("Project: %s\n", cfg.Project)
		fmt.Printf("VPS IP:  %s (WG: %s)\n", cfg.VPS.IP, cfg.VPS.WGIp)
		fmt.Printf("Home WG IP: %s\n", cfg.Home.WGIp)
		fmt.Printf("Proxy:   %s (%s)\n", cfg.Proxy.Type, cfg.Proxy.Email)
		fmt.Printf("Services (%d):\n", len(cfg.Services))
		for _, s := range cfg.Services {
			fmt.Printf("  - %s: %s -> localhost:%d\n", s.Name, s.Domain, s.Port)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
