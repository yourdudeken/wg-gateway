package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Update configuration values",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		switch strings.ToLower(key) {
		case "vps.ip":
			cfg.VPS.IP = value
		case "vps.user":
			cfg.VPS.SSHUser = value
		case "proxy.email":
			cfg.Proxy.Email = value
		case "project":
			cfg.Project = value
		default:
			fmt.Printf("Unknown config key: %s\n", key)
			return
		}

		err = config.SaveConfig("config.yaml", cfg)
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Printf("Configuration updated: %s = %s\n", key, value)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
