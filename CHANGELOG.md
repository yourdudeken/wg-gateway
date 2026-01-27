# Changelog

All notable changes to the `wg-gateway` project will be documented in this file.

## [2.0.0] - 2026-01-27
### Added
- **Multi-Node Support**: Refactored architecture to support Hub-and-Spoke topology (one VPS, many home servers).
- **Command-Driven UI**: Complete CLI lifecycle management. Users no longer need to edit files or SSH manually.
- **Local Tunnel Management**: Added `up` and `down` commands to manage Docker Compose lifecycle for peers.
- **Service Management**: Added `update` command to modify existing service mappings.
- **Security Hardening**: 
    - Automated **Fail2Ban** installation and configuration on VPS.
    - Custom SSH identity file support via `--key` and `config vps.key`.
    - UFW firewall automation.
- **Observability**:
    - New `check` command for live end-to-end connectivity auditing.
    - New `logs` command for streaming remote (VPS) and local (Home) container logs.
- **CLI Aesthetics**: Added a professional ASCII banner and removed emojis from output for enterprise compatibility.

### Changed
- Refactored `config.yaml` to support a `Peers` slice instead of a single `Home` object.
- Improved `deploy` command with better error reporting and automated SCP/SSH orchestration.
- Updated Traefik templates to dynamically resolve peer IPs based on service mapping.

## [1.0.0] - 2026-01-25
### Added
- Initial release of the `wg-gateway` CLI.
- Basic WireGuard point-to-point tunnel orchestration.
- Support for Traefik reverse proxy with automated Let's Encrypt.
- `init`, `add-service`, `generate`, and `deploy` base commands.
- Support for Docker-based deployments on both VPS and Home server.

---
*Note: This project follows Semantic Versioning.*
