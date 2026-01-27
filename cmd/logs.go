package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/yourdudeken/wg-gateway/internal/config"
)

var followFlag bool

var logsCmd = &cobra.Command{
	Use:   "logs [vps|home] [service]",
	Short: "View logs from VPS or Home server",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		service := ""
		if len(args) > 1 {
			service = args[1]
		}

		cfg, err := config.LoadConfig("config.yaml")
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if target == "vps" {
			viewVPSLogs(cfg, service)
		} else if target == "home" {
			viewHomeLogs(service)
		} else {
			fmt.Println("Invalid target. Use 'vps' or 'home'.")
		}
	},
}

func viewVPSLogs(cfg *config.Config, service string) {
	dest := fmt.Sprintf("%s@%s", cfg.VPS.SSHUser, cfg.VPS.IP)
	
	remoteCmd := "cd ~/wg-gateway && docker compose logs"
	if followFlag {
		remoteCmd += " -f"
	}
	if service != "" {
		remoteCmd += " " + service
	}

	sshArgs := []string{dest, remoteCmd}
	if cfg.VPS.SSHKey != "" {
		sshArgs = append([]string{"-i", cfg.VPS.SSHKey}, sshArgs...)
	}

	cmd := exec.Command("ssh", sshArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error viewing logs: %v\n", err)
	}
}

func viewHomeLogs(service string) {
	// For home logs, we assume the user is running this on the home server 
	// or we just run docker compose logs in the deploy/home directory
	fmt.Println("Viewing local (home) logs...")
	
	args := []string{"compose", "logs"}
	if followFlag {
		args = append(args, "-f")
	}
	if service != "" {
		args = append(args, service)
	}

	cmd := exec.Command("docker", args...)
	cmd.Dir = "deploy/home"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error viewing local logs: %v\n", err)
	}
}

func init() {
	logsCmd.Flags().BoolVarP(&followFlag, "follow", "f", false, "Follow log output")
	rootCmd.AddCommand(logsCmd)
}
