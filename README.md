```markdown
# Go Load Balancer

A simple round-robin HTTP load balancer written in Go. Incoming requests are proxied to a pool of backend servers, health-checked periodically, and automatically removed from rotation when unreachable.

---

## Features

- **Round-Robin** request distribution across healthy backends
- **Reverse proxy** using `net/http/httputil`
- **Health checks** every configurable interval to detect downed backends
- **YAML + env**-driven configuration via Viper
- Modular code structure for easy testing and extension

---

## Project Structure
```

.
├── cmd/
│ └── main.go # Entry-point: calls server.Run(":8080")
├── internal/
│ ├── configs/
│ │ └── config.go # Reads data/config.yaml with Viper
│ └── server/
│ ├── server.go # `Server` type & constructor
│ ├── pool.go # `ServerPool` & round-robin logic
│ ├── proxy.go # `Proxy` wrapper around `ReverseProxy`
│ ├── handler.go # HTTP handler that forwards to next backend
│ ├── health.go # Periodic health-check routines
│ └── run.go # `Run(addr string)` bootstraps config, pool, health, HTTP server
├── data/
│ └── config.yaml # Example configuration file
├── Makefile # Handy build & run targets
└── go.mod

````

---

## Configuration

Place your `config.yaml` under `data/`:

```yaml
server:
  host: "localhost"      # (unused by load-balancer but reserved for future use)
  port: "8080"

backends:
  - name: "backend1"
    destination_url: "http://localhost:9001"
  - name: "backend2"
    destination_url: "http://localhost:9002"
  - name: "backend3"
    destination_url: "http://localhost:9003"
````

- **`backends`**: list of backend services to proxy to.
- **`destination_url`**: full URL (scheme + host:port).

Configuration is loaded with [Viper](https://github.com/spf13/viper) .

---

## Getting Started

### Prerequisites

- Go 1.18+ installed and `$GOPATH` / modules enabled
- (Optional) `make`

### Installation

```bash
git clone https://github.com/rahulSailesh-shah/load_balancer.git
cd load_balancer
go mod download
```

### Running

#### Via Go

```bash
# Ensure you have data/config.yaml in place
go run ./cmd
```

#### Via Makefile

```bash
make build       # compiles binary to ./bin/load-balancer
make run         # builds & runs load-balancer on :8080
make backends    # (optional) runs example /ping backends on ports 9001–9003
```

---

## Example Backends

You can spin up simple HTTP backends (e.g. [in `cmd/backends.go`]) that expose `/ping` on ports 9001–9003:

```bash
go run ./cmd/backends.go
```

They’ll respond:

```
$ curl http://localhost:9001/ping
pong from server on port 9001
```

---

## How It Works

1. **Startup**: `main.go` calls `server.Run(":8080")`, which reads your YAML config .
2. **Pool Initialization**: Each backend URL is wrapped in a `Proxy` and added to `ServerPool` .
3. **Health Checks**: Every 2 minutes, each backend is probed via TCP; unreachable ones are marked down and skipped .
4. **Request Handling**: Incoming HTTP requests hit `handler.go`, which picks the next alive server (round-robin) and invokes its reverse-proxy handler .
5. **Error Handling**: If a proxy call errors, the backend is marked down and subsequent requests skip it; future retries can be implemented in `run.go`’s error handler.

---

## Extending & Testing

- Add retries or circuit-breaker logic in `run.go`’s `ErrorHandler`.
- Write unit tests for `pool.GetNextServer()`, `isAlive()`, etc., by injecting fake backends.
- Configure custom health-check intervals via a new YAML field.

---

## License

MIT © 2025 Rahul Sailesh Shah

```

```
