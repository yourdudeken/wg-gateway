package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/yourdudeken/wg-gateway/internal/config"
	"github.com/yourdudeken/wg-gateway/internal/service"
	"github.com/yourdudeken/wg-gateway/internal/wg"
)

//go:embed templates/* static/*
var content embed.FS

type Server struct {
	configPath string
}

func NewServer(configPath string) *Server {
	return &Server{configPath: configPath}
}

func (s *Server) Start(port int) error {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/api/status", s.handleStatus)
	http.HandleFunc("/api/peers", s.handlePeers)
	http.HandleFunc("/api/peers/add", s.handleAddPeer)
	http.HandleFunc("/api/services", s.handleServices)
	http.HandleFunc("/api/services/add", s.handleAddService)
	http.HandleFunc("/api/services/delete", s.handleDeleteService)
	http.HandleFunc("/api/config", s.handleConfig)
	http.HandleFunc("/api/config/update", s.handleUpdateConfig)

	http.Handle("/static/", http.FileServer(http.FS(content)))

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Web UI available at http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status := map[string]interface{}{
		"project":   cfg.Project,
		"vps_ip":    cfg.VPS.IP,
		"vps_user":  cfg.VPS.SSHUser,
		"ready":     cfg.Validate() == nil,
		"peer_count": len(cfg.Peers),
		"service_count": len(cfg.Services),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (s *Server) handlePeers(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cfg.Peers)
}

func (s *Server) handleAddPeer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
		IP   string `json:"ip"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	keys, err := wg.GenerateKeyPair()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPeer := config.PeerConfig{
		Name:       req.Name,
		WGIp:       req.IP,
		Keepalive:  25,
		PrivateKey: keys.Private,
		PublicKey:  keys.Public,
	}

	cfg.Peers = append(cfg.Peers, newPeer)

	if err := config.SaveConfig(s.configPath, cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPeer)
}

func (s *Server) handleServices(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cfg.Services)
}

func (s *Server) handleAddService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Domain   string `json:"domain"`
		Port     int    `json:"port"`
		PeerName string `json:"peer_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := service.Add(cfg, req.Domain, req.Domain, req.Port, req.PeerName); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := config.SaveConfig(s.configPath, cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (s *Server) handleDeleteService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "domain required", http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := service.Remove(cfg, domain); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := config.SaveConfig(s.configPath, cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cfg)
}

func (s *Server) handleUpdateConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cfg, err := config.LoadConfig(s.configPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch req.Key {
	case "vps.ip":
		cfg.VPS.IP = req.Value
	case "vps.user":
		cfg.VPS.SSHUser = req.Value
	case "proxy.email":
		cfg.Proxy.Email = req.Value
	case "project":
		cfg.Project = req.Value
	default:
		http.Error(w, "Unknown config key", http.StatusBadRequest)
		return
	}

	if err := config.SaveConfig(s.configPath, cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
