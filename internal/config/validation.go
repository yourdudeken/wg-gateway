package config

import (
	"errors"
)

func (c *Config) Validate() error {
	if c.Project == "" {
		return errors.New("project name cannot be empty")
	}
	if c.VPS.IP == "" {
		return errors.New("vps.ip must be set")
	}
	if c.VPS.SSHUser == "" {
		return errors.New("vps.ssh_user must be set")
	}
	if c.Proxy.Email == "" {
		return errors.New("proxy.email must be set for Let's Encrypt certificates")
	}
	return nil
}
