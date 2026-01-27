package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var addServiceCmd = &cobra.Command{
	Use:   "add-service [domain] [port]",
	Short: "Add a service to the gateway",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		port, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Invalid port: %v\n", err)
			return
		}

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		serviceName := domain // fallback to domain if not specified or cleanup domain for name
		cfg.Services = append(cfg.Services, config.Service{
			Name:   serviceName,
			Domain: domain,
			Port:   port,
		})

		err = config.SaveConfig("config.yaml", cfg)
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Printf("Service %s added on port %d\n", domain, port)
	},
}

func init() {
	rootCmd.AddCommand(addServiceCmd)
}
