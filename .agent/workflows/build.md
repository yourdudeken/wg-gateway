---
description: Build and install the wg-gateway tool locally
---
// turbo-all
1. Update Go dependencies
   ```bash
   go mod tidy
   ```
2. Build the binary
   ```bash
   go build -o wg-gateway main.go
   ```
3. Global installation (Option A: Symlink)
   ```bash
   sudo ln -sf $(pwd)/wg-gateway /usr/local/bin/wg-gateway
   ```
4. Verify the build
   ```bash
   wg-gateway hub list
   ```
