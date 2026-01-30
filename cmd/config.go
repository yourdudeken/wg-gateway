package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Update configuration values without editing files",
	Long: `Update configuration values using dot notation.
Examples:
  wg-gateway config vps.ip 1.2.3.4
  wg-gateway config vps.user root
  wg-gateway config proxy.email admin@domain.com
  wg-gateway config project my-gateway
  wg-gateway config monitor.interval 10
  wg-gateway config monitor.discord.url https://discord.com/api/webhooks/...
  wg-gateway config monitor.discord.enabled true`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		parts := strings.Split(key, ".")
		
		switch parts[0] {
		case "vps":
			if len(parts) < 2 {
				fmt.Println("Invalid key. Use vps.ip, vps.user, vps.key, etc.")
				return
			}
			switch parts[1] {
			case "ip":
				cfg.VPS.IP = value
			case "user":
				cfg.VPS.SSHUser = value
			case "key":
				cfg.VPS.SSHKey = value
			case "wg_ip":
				cfg.VPS.WGIp = value
			default:
				fmt.Printf("Unknown VPS config: %s\n", parts[1])
				return
			}
		case "proxy":
			if len(parts) < 2 {
				fmt.Println("Invalid key. Use proxy.email or proxy.type.")
				return
			}
			switch parts[1] {
			case "email":
				cfg.Proxy.Email = value
			case "type":
				cfg.Proxy.Type = value
			default:
				fmt.Printf("Unknown Proxy config: %s\n", parts[1])
				return
			}
		case "monitor":
			if len(parts) < 2 {
				fmt.Println("Invalid key. Use monitor.interval, monitor.discord.url, etc.")
				return
			}
			switch parts[1] {
			case "interval":
				fmt.Sscanf(value, "%d", &cfg.Monitor.Interval)
			case "discord":
				if len(parts) < 3 { return }
				switch parts[2] {
				case "url": cfg.Monitor.Discord.URL = value
				case "enabled": cfg.Monitor.Discord.Enabled = (value == "true")
				}
			case "telegram":
				if len(parts) < 3 { return }
				switch parts[2] {
				case "token": cfg.Monitor.Telegram.Token = value
				case "chat_id": cfg.Monitor.Telegram.ChatID = value
				case "enabled": cfg.Monitor.Telegram.Enabled = (value == "true")
				}
			}
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
