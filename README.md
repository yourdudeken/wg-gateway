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
   ./wg-gateway add-service myapp.com 8080
   ```

3. **Deploy to VPS (Automated)**
   *Requires SSH key access to your VPS.*
   ```bash
   ./wg-gateway deploy
   ```
   This will automatically:
   - Generate all configurations.
   - Upload them to your VPS via SCP.
   - Start WireGuard and Traefik on your VPS via SSH.

4. **Connect your Home Server**
   ```bash
   cd deploy/home
   # Add your services to the generated docker-compose.yaml
   docker compose up -d
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
