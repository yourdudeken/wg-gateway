package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Automate VPS setup and deployment",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.VPS.IP == "" {
			fmt.Println("Error: VPS IP is not set. Use 'wg-gateway config vps.ip <ip>' first.")
			return
		}

		// 1. Generate local files first
		fmt.Println("Generating deployment files...")
		generateFiles(cfg)

		// 2. Deploy to VPS
		fmt.Printf("Deploying to VPS (%s)...\n", cfg.VPS.IP)
		
		vpsTarget := fmt.Sprintf("%s@%s:~/wg-gateway", cfg.VPS.SSHUser, cfg.VPS.IP)
		
		// Create directory on VPS
		err = runRemoteCommand(cfg, "mkdir -p ~/wg-gateway")
		if err != nil {
			fmt.Printf("Error creating directory on VPS: %v\n", err)
			return
		}

		// SCP files to VPS
		fmt.Println("Uploading files...")
		err = runLocalCommand("scp", "-r", "deploy/vps/.", vpsTarget)
		if err != nil {
			fmt.Printf("Error uploading files to VPS: %v\n", err)
			return
		}

		// Run docker-compose on VPS
		fmt.Println("Starting services on VPS...")
		err = runRemoteCommand(cfg, "cd ~/wg-gateway && docker compose up -d || docker-compose up -d")
		if err != nil {
			fmt.Printf("Error starting services on VPS: %v\n", err)
			return
		}

		fmt.Println("\nSuccess! VPS is set up and tunnel is live.")
		fmt.Println("Now run 'docker compose up -d' in your 'deploy/home' directory to connect your home server.")
	},
}

func runLocalCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runRemoteCommand(cfg *config.Config, remoteCmd string) error {
	dest := fmt.Sprintf("%s@%s", cfg.VPS.SSHUser, cfg.VPS.IP)
	cmd := exec.Command("ssh", dest, remoteCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func generateFiles(cfg *config.Config) {
	// Call the same logic as generateCmd
	generateAction(cfg)
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
