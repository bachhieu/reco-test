# Test

A new project created with fkit make:init command.

## Getting Started

1. Clone the repository
2. Run the application:
   ```bash
   cp example.config.json config.json
   go run main.go
   ```

## Project Structure

- biz/: Business logic
  - core/: Core business logic
  - dal/: Data access layer
    - dao/: Data access objects
    - do/: Data object
    - models/: Models
- cmd/: Job commands
- docs/: Documentation
- pkg/: Shared packages
- scripts/: Utility scripts
- server/: HTTP server implementation
