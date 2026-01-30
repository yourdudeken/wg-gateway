package service

type AppTemplate struct {
	Name string
	Port int
}

var Templates = map[string]AppTemplate{
	"plex": {
		Name: "Plex Media Server",
		Port: 32400,
	},
	"ha": {
		Name: "Home Assistant",
		Port: 8123,
	},
	"home-assistant": {
		Name: "Home Assistant",
		Port: 8123,
	},
	"jellyfin": {
		Name: "Jellyfin",
		Port: 8096,
	},
	"pihole": {
		Name: "Pi-hole",
		Port: 80,
	},
	"grafana": {
		Name: "Grafana",
		Port: 3000,
	},
	"prom": {
		Name: "Prometheus",
		Port: 9090,
	},
	"prometheus": {
		Name: "Prometheus",
		Port: 9090,
	},
	"overseerr": {
		Name: "Overseerr",
		Port: 5055,
	},
	"tautulli": {
		Name: "Tautulli",
		Port: 8181,
	},
}
