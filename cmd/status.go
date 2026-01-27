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
			fmt.Printf("❌ Error loading config: %v\n", err)
			return
		}

		fmt.Println("📊 W-G Gateway Status")
		fmt.Println("======================")
		fmt.Printf("Project:    %s\n", cfg.Project)
		fmt.Printf("VPS IP:     %s (User: %s)\n", cfg.VPS.IP, cfg.VPS.SSHUser)
		fmt.Printf("Audit:      ")
		if err := cfg.Validate(); err != nil {
			fmt.Printf("⚠️  Incomplete: %v\n", err)
		} else {
			fmt.Println("✅ Production Ready")
		}
		
		fmt.Println("\n🌐 Networking")
		fmt.Printf("  Tunnel IP (VPS):  %s\n", cfg.VPS.WGIp)
		fmt.Printf("  Tunnel IP (Home): %s\n", cfg.Home.WGIp)
		
		fmt.Println("\n🛠️  Services")
		if len(cfg.Services) == 0 {
			fmt.Println("  No services configured.")
		} else {
			for i, s := range cfg.Services {
				fmt.Printf("  %d. %s -> 10.0.0.2:%d\n", i+1, s.Domain, s.Port)
			}
		}
		fmt.Println("======================")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
