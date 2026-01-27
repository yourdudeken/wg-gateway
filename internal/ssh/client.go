package ssh

import (
	"fmt"
	"os"
	"os/exec"
)

type Client struct {
	User string
	Host string
}

func NewClient(user, host string) *Client {
	return &Client{User: user, Host: host}
}

func (c *Client) Run(command string) error {
	dest := fmt.Sprintf("%s@%s", c.User, c.Host)
	cmd := exec.Command("ssh", "-o", "BatchMode=yes", "-o", "ConnectTimeout=10", dest, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Client) Copy(src, dst string) error {
	dest := fmt.Sprintf("%s@%s:%s", c.User, c.Host, dst)
	cmd := exec.Command("scp", "-o", "BatchMode=yes", "-r", src, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
