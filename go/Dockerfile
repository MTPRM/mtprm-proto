FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ cmd/rest.go


FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/bin/rest /opt/app/rest

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/opt/app/rest" ]
