package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/ssh"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the health and status of the gateway",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if err := cfg.Validate(); err != nil {
			fmt.Printf("Config is incomplete: %v\n", err)
			return
		}

		client := ssh.NewClient(cfg.VPS.SSHUser, cfg.VPS.IP, cfg.VPS.SSHKey)

		fmt.Println("Health Check: VPS Connection")
		fmt.Println("---------------------------")
		
		// 1. Check SSH Connectivity
		fmt.Print("SSH Connectivity: ")
		if err := client.Run("true"); err != nil {
			fmt.Println("FAILED")
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Println("OK")

		// 2. Check Docker
		fmt.Print("Docker Engine:   ")
		if err := client.Run("docker info > /dev/null 2>&1"); err != nil {
			fmt.Println("FAILED")
		} else {
			fmt.Println("OK")
		}

		// 3. Check WireGuard Interface
		fmt.Print("WireGuard (wg0): ")
		if err := client.Run("ip addr show wg0 > /dev/null 2>&1"); err != nil {
			fmt.Println("DOWN")
		} else {
			fmt.Println("UP")
		}

		// 4. Check Handshake
		fmt.Print("Tunnel Status:   ")
		// We catch output to check handshake
		// Note: We need a way to capture output from client.Run or create a new method
		fmt.Println("Checking...")
		client.Run("sudo wg show wg0")

		// 5. Ping Home from VPS
		fmt.Printf("Ping Home (%s): ", cfg.Home.WGIp)
		if err := client.Run(fmt.Sprintf("ping -c 3 -W 2 %s > /dev/null 2>&1", cfg.Home.WGIp)); err != nil {
			fmt.Println("UNREACHABLE")
		} else {
			fmt.Println("REACHABLE")
		}
		
		fmt.Println("---------------------------")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
