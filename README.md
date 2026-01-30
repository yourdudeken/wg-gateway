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

---

## Core Philosophy: Zero Manual Configuration

*   **No SSH Required**: The tool manages the VPS via automated orchestration.
*   **No File Editing**: All configurations are handled via the CLI.
*   **Hub-and-Spoke**: One VPS can securely tunnel to multiple distributed home server nodes.
*   **Zero-Setup DNS**: Built-in `sslip.io` support for those without a domain.
*   **Security First**: Automated hardening with Fail2Ban, UFW, and custom SSH identity support.

---

## Feature Highlights

*   **Automated VPS Provisioning**: One command to transform a fresh Ubuntu VPS into a secured gateway.
*   **Multi-Node Support**: Connect multiple home servers (peers) to a single VPS hub.
*   **Observability Suite**: Real-time health checks, handshakes, and container log streaming.
*   **Docker Native**: Runs entirely in containers for isolation and easy cleanup.
*   **Auto-HTTPS**: Integrated Let's Encrypt support via Traefik.
*   **Emoji-Free**: Clean, professional interface for developers.

---

---

## Installation Options

Choose the method that best fits your environment.

### 1. The One-Liner (Recommended for most users)
This script automatically detects your OS and architecture, downloads the latest binary, and installs it to your system.
```bash
curl -sSfL https://raw.githubusercontent.com/yourdudeken/wg-gateway/main/scripts/install.sh | sh
```

### 2. For Go Developers
If you have Go installed, you can install the tool directly from source:
```bash
go install github.com/yourdudeken/wg-gateway@latest
```

### 3. Manual Build (From Source)
If you prefer to build it manually:
```bash
git clone https://github.com/yourdudeken/wg-gateway.git
cd wg-gateway

go build -o wg-gateway main.go
sudo ln -sf $(pwd)/wg-gateway /usr/local/bin/wg-gateway
```

---

## Quick Start (5-Step Workflow)

### 1. Initialize & Setup
Initialize your project and run the setup command to configure your local firewall (UFW) to allow tunnel traffic.
```bash
wg-gateway init --ip 1.2.3.4 --user root --key ~/.ssh/id_ed25519 --email admin@example.com
wg-gateway setup
```

### 2. Manage Peers (Nodes)
Add your home server nodes. Each peer receives a unique WireGuard configuration.
```bash
wg-gateway peer add warehouse-lab
```

### 3. Add Services (Zero-Setup DNS)
Route domains to your peers. If you do not have a domain, use the built-in sslip.io support. Providing a hostname without a dot automatically maps it to your VPS IP.
```bash
wg-gateway service add dash 8080 --peer home
# Result: dash.1.2.3.4.sslip.io (No DNS configuration required)
```

### 4. Deploy to VPS
Bootstrap the remote server and sync configurations.
```bash
wg-gateway deploy --bootstrap
```

### 5. Start the Tunnel
Launch the secure tunnel on your local machine.
```bash
wg-gateway up home
```

---

## Web UI Dashboard

For users who prefer a graphical interface, `wg-gateway` includes a modern web-based dashboard:

```bash
wg-gateway web --port 8080 --password MySecurePass
```
*Note: You can also set the `WG_ADMIN_PASS` environment variable instead of using the flag.*

Once started, the dashboard is available at `http://localhost:8080` (Username: `admin`).

### Security Note
If a password is provided, the dashboard is protected by Basic Authentication. If no password is provided, the dashboard is public—only use this mode in trusted local environments.

---

## Command Reference

### System & Local Setup
*   `setup`: Configure the local firewall (UFW) and verify system readiness.
*   `init`: Create a new project configuration.

### Infrastructure Lifecycle
*   `deploy [--bootstrap]`: Provision VPS, install Docker/WireGuard, and upload configs.
*   `generate`: Manually render the deployment Docker and WireGuard files.
*   `destroy`: Wipe the local `deploy/` directory.

### Node & Service Management
*   `peer add [name]`: Register a new home server node.
*   `peer list`: View all configured nodes and their tunnel IPs.
*   `service add [domain] [port]`: Map a domain to a local port (Supports sslip.io).
*   `service update [domain] [port]`: Change the target port for a domain.
*   `service list`: View all active routing rules.

### Local Execution
*   `up [peer]`: Start the WireGuard tunnel for a specific peer.
*   `down [peer]`: Stop the tunnel for a specific peer.
*   `web [--port]`: Launch the web-based dashboard interface.

### Observability & Maintenance
*   `status`: Overview of the project, including a production-readiness audit.
*   `check`: Live connectivity test (pings all peers and checks service ports).
*   `logs [vps|peer_name]`: Stream real-time logs from remote or local containers.
*   `config [key] [value]`: Update any setting (e.g., `vps.ip`, `proxy.email`) via CLI.
*   `rotate-keys`: Regenerate all WireGuard keypairs for the hub and spokes.

---

## Security

`wg-gateway` implements several hardening measures automatically:
*   **Fail2Ban**: Installed and configured during bootstrap to block brute-force attacks.
*   **UFW Firewall**: Defaults to "deny all" with explicit allows for SSH (22), WireGuard (51820), and Web (80/443).
*   **Custom SSH Keys**: Supports specific identity files for deployment.
*   **Internal Networking**: Services are bridged via a private WireGuard network, never exposed directly on the VPS host.

---

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
