package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/wg"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := "config.yaml"
		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("Config file already exists.")
			return
		}

		cfg := config.NewDefaultConfig()

		vpsKeys, err := wg.GenerateKeyPair()
		if err != nil {
			fmt.Printf("Error generating VPS keys: %v\n", err)
			return
		}
		cfg.VPS.PrivateKey = vpsKeys.Private
		cfg.VPS.PublicKey = vpsKeys.Public

		homeKeys, err := wg.GenerateKeyPair()
		if err != nil {
			fmt.Printf("Error generating Home keys: %v\n", err)
			return
		}
		cfg.Home.PrivateKey = homeKeys.Private
		cfg.Home.PublicKey = homeKeys.Public

		err = config.SaveConfig(configPath, cfg)
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Println("Project initialized. Please edit config.yaml to set your VPS IP and services.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
