package config

type Config struct {
	Project string   `yaml:"project"`
	VPS     VPSConfig `yaml:"vps"`
	Home    HomeConfig `yaml:"home"`
	Proxy   ProxyConfig `yaml:"proxy"`
	Services []Service `yaml:"services"`
}

type VPSConfig struct {
	IP         string `yaml:"ip"`
	SSHUser    string `yaml:"ssh_user"`
	WGIp       string `yaml:"wg_ip"`
	WGPort     int    `yaml:"wg_port"`
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
}

type HomeConfig struct {
	WGIp       string `yaml:"wg_ip"`
	Keepalive  int    `yaml:"keepalive"`
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
}

type ProxyConfig struct {
	Type  string `yaml:"type"` // e.g. "traefik"
	Email string `yaml:"email"`
}

type Service struct {
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
	Port   int    `yaml:"port"`
}
