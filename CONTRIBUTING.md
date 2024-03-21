# Getting Started

## Requirements
1. pnpm
2. golang

```bash
go run cmd/maestro/maestro.go
pnpm run watch
```

## Local setup

### Default address
```bash
go run main.go serve # Default to localhost:8080
```
### Custom address
```bash
export MAESTRO_LISTEN_ADDRESS="0.0.0.0:9000"
go run main.go serve
```

### Functional tests
```bash
go run main.go conduct cpu saturation 50.0 5s
```