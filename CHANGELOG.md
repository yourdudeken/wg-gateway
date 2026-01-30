# Changelog

All notable changes to the `wg-gateway` project will be documented in this file.

## [2.2.0] - 2026-01-31
### Added
- **Web UI Authentication (Basic Auth)**: Secure the dashboard with a password using the `--password` flag or `WG_ADMIN_PASS` environment variable.
- **Zero-Setup DNS (sslip.io)**: Automatically map hostnames to the VPS IP for users without custom domains.
- **Local Setup Engine**: New `setup` command to automate local environment configuration.
    - Automated UFW firewall rule injection (`sudo ufw allow in on wg0`)
    - Docker daemon health verification
    - Global installation path generator
- **Enhanced Log Filtering**: Improved `logs` command to support peer-specific log streaming using the new distributed generation architecture.
- **Improved Health Auditing**: The `check` command now performs deep-packet verification of service port connectivity from the VPS hub.

## [2.1.0] - 2026-01-27
### Added
- **Web UI Dashboard**: Modern, dark-themed web interface for visual gateway management.
    - Real-time status monitoring with live statistics
    - Visual peer and service management
    - Configuration editor with form validation
    - Responsive design for desktop and mobile
    - RESTful API backend with embedded static files
- **Enhanced CLI Commands**:
    - `web`: Launch the web-based dashboard interface
    - `up [peer]`: Start local peer tunnel via Docker Compose
    - `down [peer]`: Stop local peer tunnel
    - `service update`: Modify existing service port mappings
- **Improved Configuration Management**:
    - Dot-notation support in `config` command (e.g., `vps.ip`, `proxy.email`)
    - Enhanced validation and error messages

### Changed
- Removed all emojis from codebase for enterprise compatibility
- Updated README with comprehensive command reference and web UI documentation
- Improved CLI help text and command descriptions

## [2.0.0] - 2026-01-27
### Added
- **Multi-Node Support**: Refactored architecture to support Hub-and-Spoke topology (one VPS, many home servers).
- **Command-Driven UI**: Complete CLI lifecycle management. Users no longer need to edit files or SSH manually.
- **Peer Management**: 
    - New `peer` command group for managing multiple home servers
    - Automatic IP assignment for new peers
    - Key generation and rotation for all peers
- **Service Management**: 
    - `service add` with `--peer` flag for multi-node routing
    - `service remove` for cleanup
    - `service list` with peer association display
- **Security Hardening**: 
    - Automated **Fail2Ban** installation and configuration on VPS
    - Custom SSH identity file support via `--key` and `config vps.key`
    - UFW firewall automation with explicit port rules
- **Observability Suite**:
    - `check` command for live end-to-end connectivity auditing
    - `logs` command for streaming remote (VPS) and local (Home) container logs
    - Enhanced `status` command with production-readiness audit
- **CLI Aesthetics**: 
    - Professional ASCII banner
    - Clean, emoji-free output for enterprise logs

### Changed
- Refactored `config.yaml` to support a `Peers` slice instead of a single `Home` object
- Updated all WireGuard and Traefik templates for multi-peer support
- Improved `deploy` command with better error reporting and automated SCP/SSH orchestration
- Enhanced `generate` command to create peer-specific deployment folders

### Breaking Changes
- Configuration file structure changed from `Home` to `Peers` array
- Service definitions now require `peer_name` field
- Deployment files now generated in `deploy/peers/[name]` instead of `deploy/home`

## [1.0.0] - 2026-01-25
### Added
- Initial release of the `wg-gateway` CLI
- Basic WireGuard point-to-point tunnel orchestration
- Support for Traefik reverse proxy with automated Let's Encrypt
- Core commands: `init`, `generate`, `deploy`, `status`
- Docker-based deployments for both VPS and Home server
- Automated WireGuard key generation
- Basic service routing configuration

---
*Note: This project follows Semantic Versioning.*
