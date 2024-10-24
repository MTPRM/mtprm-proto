# grpc-gateway

## Manual Testing

`go mod download`
`go build -o bin/ cmd/rest.go`

`./bin/rest ${grpcServerBackend}:${port}`
`./bin/rest api-portfolio.whitehawk-mtprm.com:443`

`firefox localhost:8080`
