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
   go run main.go init
   ```
   This creates a `config.yaml` with generated WireGuard keys.

2. **Configure your VPS**
   Edit `config.yaml` and set:
   - `vps.ip`: Your VPS public IP address.
   - `proxy.email`: Your email for Let's Encrypt certificates.

3. **Add a service**
   ```bash
   go run main.go add-service api.example.com 3000
   ```

4. **Generate deployment files**
   ```bash
   go run main.go generate
   ```
   This creates the `deploy/` directory.

5. **Deploy to VPS**
   Copy `deploy/vps` to your VPS and run:
   ```bash
   docker-compose up -d
   ```

6. **Deploy to Home Server**
   Copy `deploy/home` to your home server.
   Add your services to the `docker-compose.yaml` using `network_mode: "service:wireguard"`.
   Then run:
   ```bash
   docker-compose up -d
   ```

## Workflow
1. User → `api.example.com` (VPS)
2. Traefik (VPS) → WireGuard Tunnel (`10.0.0.2`)
3. Home Server Service (`10.0.0.2:3000`)

## Commands
- `init`: Initialize a new project and generate keys.
- `add-service [domain] [port]`: Register a new home service.
- `generate`: Render Docker and WireGuard configurations.
- `status`: Show current configuration.
- `rotate-keys`: Regenerate all WireGuard keys.
- `destroy`: Remove generated deployment files.

## Contributing
Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to get involved.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
