package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/service"
)

var (
	targetPeer string
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage home services",
}

var addSvcCmd = &cobra.Command{
	Use:   "add [domain] [port]",
	Short: "Add a new service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		port, _ := strconv.Atoi(args[1])

		if err := service.Validate(domain, port); err != nil {
			fmt.Printf("Validation error: %v\n", err)
			return
		}

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		// Use domain as name if not specified
		name := domain

		if err := service.Add(cfg, name, domain, port, targetPeer); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if err := config.SaveConfig("config.yaml", cfg); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Printf("Service %s added successfully to peer %s.\n", domain, targetPeer)
	},
}

// ... rest of commands ...

var removeSvcCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove an existing service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if err := service.Remove(cfg, domain); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if err := config.SaveConfig("config.yaml", cfg); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Printf("Service %s removed successfully.\n", domain)
	},
}

var listSvcCmd = &cobra.Command{
	Use:   "list",
	Short: "List all services",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if len(cfg.Services) == 0 {
			fmt.Println("No services configured.")
			return
		}

		fmt.Println("Current Services:")
		for _, s := range cfg.Services {
			fmt.Printf("  - %s -> %s:%d (Peer: %s)\n", s.Domain, "localhost", s.Port, s.PeerName)
		}
	},
}

func init() {
	addSvcCmd.Flags().StringVar(&targetPeer, "peer", "home", "Target peer name for the service")
	serviceCmd.AddCommand(addSvcCmd)
	serviceCmd.AddCommand(removeSvcCmd)
	serviceCmd.AddCommand(listSvcCmd)
	rootCmd.AddCommand(serviceCmd)
}
