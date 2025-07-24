# Garantex USDT Rates GRPC Service

A Go microservice that fetches the USDT exchange rate (ask/bid) from the Grinex exchange and exposes it via gRPC.  
Exchange data is stored in PostgreSQL on every request.

---

## Features

- gRPC endpoint `GetRates` for retrieving USDT ask/bid prices
- Data fetched from [Grinex API](https://grinex.io/api/v2/depth)
- PostgreSQL integration
- Healthcheck endpoint
- Graceful shutdown
- Configurable via `.yaml` config and environment variables
- Dockerized with Docker Compose
- Linting and unit testing support

---

## 🏁 Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/Webblurt/garantex-usdt-rates.git
cd garantex-usdt-rates
```

### 2. Create `.env` file

```bash
cat <<EOF > .env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=rates
EOF
```

### 3. Build the project

```bash
make build
```

### 4. Start services with Docker Compose

```bash
docker-compose up --build
```

The app will be available on `localhost:50051` (gRPC port).

---

## Running Tests

```bash
make test
```

---

## Configuration

Configuration is managed via a YAML file and environment variables.

### config/config.yaml

```yaml
server:
  port: "50051"

postgres:
  dsn: "host=postgres port=5432 user=postgres password=postgres dbname=rates sslmode=disable"

grinex:
  url: "https://grinex.io/api/v2/depth"
```

### Environment Variables (`.env`)

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=rates
```

You can override values in the config file using environment variables or command-line flags (if implemented).

---

## Makefile Commands

`make build` - Build the Go app
`make run` - Run the app locally
`make test` - Run unit tests
`make lint` - Run code linter (`golangci-lint`)
`make docker-build` - Build Docker image

---

##  Healthcheck

A gRPC health check endpoint is available (standard [grpc-health-probe](https://github.com/grpc-ecosystem/grpc-health-probe) compatible).

---

## Requirements

- Go 1.22+
- Docker & Docker Compose
- `golangci-lint` (for `make lint`)

---

##  Running without Docker

```bash
go run ./cmd/app
```

Ensure PostgreSQL is running locally and credentials match `config/config.yaml`.

---

##  GRPC API

### GetRates (Unary)

Returns current USDT ask and bid prices with timestamp.

```proto
rpc GetRates (google.protobuf.Empty) returns (RateResponse);
```
