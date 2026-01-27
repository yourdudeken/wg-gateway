package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/templates"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate deployment files",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.VPS.IP == "" {
			fmt.Println("Warning: vps.ip is not set in config.yaml. Please set it before deploying.")
		}

		// Cleanup old deploy dir
		os.RemoveAll("deploy")

		// VPS Generation
		vpsDir := "deploy/vps"
		os.MkdirAll(filepath.Join(vpsDir, "wireguard"), 0755)
		os.MkdirAll(filepath.Join(vpsDir, "traefik_dynamic"), 0755)
		os.MkdirAll(filepath.Join(vpsDir, "letsencrypt"), 0755)

		renderAndSave(vpsDir, "docker-compose.vps.yaml.tmpl", "docker-compose.yaml", cfg)
		renderAndSave(filepath.Join(vpsDir, "wireguard"), "wg0-server.conf.tmpl", "wg0.conf", cfg)
		renderAndSave(filepath.Join(vpsDir, "traefik_dynamic"), "traefik_dynamic.yaml.tmpl", "dynamic.yaml", cfg)

		// Home Generation
		homeDir := "deploy/home"
		os.MkdirAll(filepath.Join(homeDir, "wireguard"), 0755)

		renderAndSave(homeDir, "docker-compose.home.yaml.tmpl", "docker-compose.yaml", cfg)
		renderAndSave(filepath.Join(homeDir, "wireguard"), "wg0-client.conf.tmpl", "wg0.conf", cfg)

		fmt.Println("Success! Deployment files generated in 'deploy/'")
		fmt.Println("\nTo deploy to VPS:")
		fmt.Printf("  scp -r deploy/vps %s@%s:~/wg-gateway\n", cfg.VPS.SSHUser, cfg.VPS.IP)
		fmt.Println("  ssh into VPS and run: cd ~/wg-gateway && docker-compose up -d")
		
		fmt.Println("\nTo deploy to Home Server:")
		fmt.Println("  Copy deploy/home to your home server.")
		fmt.Println("  Add your apps to deploy/home/docker-compose.yaml using 'network_mode: service:wireguard'")
		fmt.Println("  Run: docker-compose up -d")
	},
}

func renderAndSave(dir, tmplName, fileName string, cfg *config.Config) {
	data, err := templates.Render(tmplName, cfg)
	if err != nil {
		fmt.Printf("Error rendering %s: %v\n", tmplName, err)
		return
	}

	err = os.WriteFile(filepath.Join(dir, fileName), data, 0644)
	if err != nil {
		fmt.Printf("Error saving %s: %v\n", fileName, err)
		return
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
