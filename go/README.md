# grpc-gateway

## Manual Testing

### Download dependencies

`go mod download`

### Build binary

`go build -o bin/ cmd/rest.go`

### Run grpc-gateway locally

The first argument is `${grpcServerBackend}:${port}`.

E.G.

- `./bin/rest localhost:50051` to run against local GRPC server
- `./bin/rest api-portfolio.whitehawk-mtprm.com:443` to run against remote GRPC server

### Open documentation

`firefox localhost:8080`
