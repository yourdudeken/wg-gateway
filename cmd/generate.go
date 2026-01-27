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

		if err := cfg.Validate(); err != nil {
			fmt.Printf("Config Validation Warning: %v\n", err)
		}

		generateAction(cfg)
	},
}

func generateAction(cfg *config.Config) {
	// Cleanup old deploy dir
	os.RemoveAll("deploy")

	// VPS Generation
	vpsDir := "deploy/vps"
	os.MkdirAll(filepath.Join(vpsDir, "wireguard"), 0755)
	os.MkdirAll(filepath.Join(vpsDir, "traefik_dynamic"), 0755)
	os.MkdirAll(filepath.Join(vpsDir, "letsencrypt"), 0755)

	if err := renderAndSave(vpsDir, "docker-compose.vps.yaml.tmpl", "docker-compose.yaml", cfg); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := renderAndSave(filepath.Join(vpsDir, "wireguard"), "wg0-server.conf.tmpl", "wg0.conf", cfg); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := renderAndSave(filepath.Join(vpsDir, "traefik_dynamic"), "traefik_dynamic.yaml.tmpl", "dynamic.yaml", cfg); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Home Generation
	homeDir := "deploy/home"
	os.MkdirAll(filepath.Join(homeDir, "wireguard"), 0755)

	if err := renderAndSave(homeDir, "docker-compose.home.yaml.tmpl", "docker-compose.yaml", cfg); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := renderAndSave(filepath.Join(homeDir, "wireguard"), "wg0-client.conf.tmpl", "wg0.conf", cfg); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Success! Deployment files generated in 'deploy/'")
}

func renderAndSave(dir, tmplName, fileName string, cfg *config.Config) error {
	data, err := templates.Render(tmplName, cfg)
	if err != nil {
		return fmt.Errorf("rendering %s: %w", tmplName, err)
	}

	err = os.WriteFile(filepath.Join(dir, fileName), data, 0644)
	if err != nil {
		return fmt.Errorf("saving %s: %w", fileName, err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
