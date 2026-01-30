# W-G Gateway

```
 __      __  ________        ________         __                                           
/  \    /  |/  _____/       /  _____/  ____ _/  |_  ____ __  _  __ _____  ___.__.          
\   \/\/   /   \  ___      /   \  ____/ __ \\   __\/ __ \\ \/ \/ // __ \<   |  |          
 \        /    \    \     \    \_\  \  ___/ |  | \  ___/ \     /\  ___/ \___  |          
  \__/\  / \______  /      \______  /\___  >|__|  \___  > \/\_/  \___  >/ ____|          
       \/         \/               \/     \/           \/             \/ \/               
```

**WireGuard VPS-to-Home Gateway Tool**

`wg-gateway` is a production-grade infrastructure tool designed to securely expose home servers (behind CGNAT, 4G, or dynamic IPs) to the public internet. It automates the orchestration of WireGuard tunnels and Traefik reverse proxies using a strictly **command-driven workflow**.

## Core Philosophy: Zero Manual Configuration
- **No SSH Required**: The tool manages the VPS via automated orchestration.
- **No File Editing**: All configurations are handled via the CLI.
- **Hub-and-Spoke**: One VPS can securely tunnel to multiple distributed home server nodes.
- **Security First**: Automated hardening with Fail2Ban, UFW, and custom SSH identity support.

## Feature highlights
- **Automated VPS Provisioning**: One command to transform a fresh Ubuntu VPS into a secured gateway.
- **Multi-Node Support**: Connect multiple home servers (peers) to a single VPS hub.
- **Observability Suite**: Real-time health checks, handshakes, and container log streaming.
- **Docker Native**: Runs entirely in containers for isolation and easy cleanup.
- **Auto-HTTPS**: Integrated Let's Encrypt support via Traefik.

---

## Quick Start (5-Step Workflow)

### 1. Build & Initialize
Build the CLI tool and initialize your project with VPS details.
```bash
go build -o wg-gateway main.go

./wg-gateway init --ip 1.2.3.4 --user root --key ~/.ssh/id_ed25519 --email admin@example.com
```

### 2. Manage Peers (Nodes)
Add your home server nodes. The first peer 'home' is created by default.
```bash
./wg-gateway peer add warehouse-lab
```

### 3. Add Services
Route public domains to specific ports on your peers.
```bash
./wg-gateway service add myapp.com 8080 --peer home
./wg-gateway service add inventory.net 9000 --peer warehouse-lab
```

### 4. Deploy to VPS
Bootstrap the remote server (first time only) and sync configurations.
```bash
./wg-gateway deploy --bootstrap
```

### 5. Start the Tunnel
Launch the secure tunnel on your local machine.
```bash
./wg-gateway up home
```

---

## Web UI Dashboard

For users who prefer a graphical interface, `wg-gateway` includes a modern web-based dashboard:

```bash
./wg-gateway web --port 8080
```

Then open `http://localhost:8080` in your browser.

### Dashboard Features
- Real-time status monitoring with live statistics
- Visual peer and service management
- Configuration editor with form validation
- Dark-themed, responsive design
- No emojis - clean professional interface

---

## Command Reference

### Infrastructure Lifecycle
- `init`: Create a new project configuration.
- `deploy [--bootstrap]`: Provision VPS, install Docker/WireGuard, and upload configs.
- `generate`: Manually render the deployment Docker and WireGuard files.
- `destroy`: Wipe the local `deploy/` directory.

### Node & Service Management
- `peer add [name]`: Register a new home server node.
- `peer list`: View all configured nodes and their tunnel IPs.
- `service add [domain] [port]`: Map a domain to a local port.
- `service update [domain] [port]`: Change the target port for a domain.
- `service list`: View all active routing rules.

### Local Execution
- `up [peer]`: Start the WireGuard tunnel for a specific peer.
- `down [peer]`: Stop the tunnel for a specific peer.
- `web [--port]`: Launch the web-based dashboard interface.

### Observability & Maintenance
- `status`: Overview of the project, including a production-readiness audit.
- `check`: Live connectivity test (pings all peers from the VPS).
- `logs [vps|home] [service]`: Stream real-time logs from remote or local containers.
- `config [key] [value]`: Update any setting (e.g., `vps.ip`, `proxy.email`) via CLI.
- `rotate-keys`: Regenerate all WireGuard keypairs for the hub and spokes.

---

## Security
`wg-gateway` implements several hardening measures automatically:
- **Fail2Ban**: Installed and configured during bootstrap to block brute-force attacks.
- **UFW Firewall**: Defaults to "deny all" with explicit allows for SSH (22), WireGuard (51820), and Web (80/443).
- **Custom SSH Keys**: Supports specific identity files for deployment.
- **No Emojis**: Clean, professional console output suitable for enterprise logs.

## Contributing
Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
