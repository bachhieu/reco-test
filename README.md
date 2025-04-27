# Test

A new project created with fkit make:init command.

## Getting Started

1. Clone the repository and run:

step 1: Clone github repo
  ```bash
  git clone https://github.com/bachhieu/reco-test
  cp example.config.json config.json #Edit config postgres and redis 
  go run main.go
  ```   
OR
  ```bash
   cp example.config.json config.json
   docker compose up -d
   ```

2. Run with docker image: [bachviethieu/demo](https://hub.docker.com/r/bachviethieu/demo)

  ```bash
  docker pull bachviethieu/demo
  # Create file config.json
  docker run -p {PORT}:{PORT} -v ./config.json:/app/config.json bachviethieu/demo
  ```

## Testing API in local

Use swagger to testing api in local `http://localhost:{PORT}/api/v1/admin/api-docs/index.html`



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
