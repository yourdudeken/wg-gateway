package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/wg"
)

var rotateKeysCmd = &cobra.Command{
	Use:   "rotate-keys",
	Short: "Rotate WireGuard keys",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		vpsKeys, _ := wg.GenerateKeyPair()
		cfg.VPS.PrivateKey = vpsKeys.Private
		cfg.VPS.PublicKey = vpsKeys.Public

		homeKeys, _ := wg.GenerateKeyPair()
		cfg.Home.PrivateKey = homeKeys.Private
		cfg.Home.PublicKey = homeKeys.Public

		err = config.SaveConfig("config.yaml", cfg)
		if err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}

		fmt.Println("Keys rotated successfully. Run 'generate' to update deployment files.")
	},
}

func init() {
	rootCmd.AddCommand(rotateKeysCmd)
}
