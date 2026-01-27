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
3. Install to your local bin (optional)
   ```bash
   sudo mv wg-gateway /usr/local/bin/
   ```
