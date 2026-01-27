# WireGuard VPS-to-Home Gateway Tool (`wg-gateway`)

Automate exposing your home server behind CGNAT/4G through a VPS with a public IP.

## Features
- **WireGuard Automation**: Automatic key generation and config management.
- **Docker-Native**: Runs everything in Docker for no host pollution.
- **Traefik Integration**: Automatic HTTPS and host-based routing.
- **Persistent Connection**: Built-in keepalives for stable 4G/mobile connections.

## Quick Start

1. **Initialize the project**
   ```bash
   # Build the tool first
   go build -o wg-gateway main.go

   # Initialize with VPS details
   ./wg-gateway init --ip 1.2.3.4 --user root --email admin@example.com
   ```

2. **Add your services**
   ```bash
   ./wg-gateway service add myapp.com 8080
   ```

3. **Deploy to VPS (Automated)**
   *Requires SSH key access to your VPS.*
   ```bash
   # Use --bootstrap to install Docker, WireGuard, and setup Firewall on a new VPS
   ./wg-gateway deploy --bootstrap
   ```
   This will automatically:
   - **Bootstrap**: Update system, install Docker/WireGuard, and configure UFW Firewall.
   - **Provision**: Generate configurations and upload them to your VPS.
   - **Launch**: Start WireGuard and Traefik containers on the remote host.

4. **Connect your Home Server**
   ```bash
   cd deploy/home
   # Add your services to the generated docker-compose.yaml
   docker compose up -d
   ```

## Peer Management (Multi-Node)
The tool supports a **Hub-and-Spoke** topology, allowing one VPS to serve multiple home servers.
- `peer add [name]`: Add a new home server peer.
- `peer list`: View all configured peers/branches.

Example:
```bash
./wg-gateway peer add warehouse-lab
./wg-gateway service add lab-api.com 8080 --peer warehouse-lab
```

## Service Management
Manage your home services effortlessly via CLI:
- `add [domain] [port] [--peer name]`: Add a new routing rule (defaults to 'home').
- `remove [domain]`: Delete an existing service.
- `list`: View all active services and their associated peers.

Example:
```bash
./wg-gateway service add api.example.com 3000
./wg-gateway service remove old.example.com
```

## Observability
Keep an eye on your gateway's health with built-in monitoring:
- `check`: Connectivity test for VPS and **all** connected peers.
- `logs vps traefik`: Streams logs from the Traefik proxy on your VPS.
- `logs home`: Views logs for your 'home' peer WireGuard container.

## Commands
- `init`: Initialize project with defaults or flags.
- `peer`: Manage home server peers (add, list).
- `service`: Manage home services (add, remove, list).
- `deploy`: One-click deployment to VPS (with optional `--bootstrap`).
- `config`: Update specific configuration keys.
- `status`: Health check and configuration overview.
- `check`: Perform a live connectivity and tunnel health check.
- `logs [vps|home]`: View or follow container logs.
- `generate`: Manually render configuration files.
- `rotate-keys`: Securely cycle WireGuard keys.
- `destroy`: Clean up 'deploy' directory.

## Contributing
Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to get involved.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
