package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	entities_api "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities/v1"
	entities__0__reports__combined__xlsx "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities__0__reports__combined__xlsx/v1"
	entities__0__reports__summary "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities__0__reports__summary/v1"
	reports__combined__zip "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/reports__combined__zip/v1"
)

func main() {
	port := ":8080"

	log.Println("Starting listening on port:", port)
	mux := runtime.NewServeMux()

	grpcServerEndpoint := os.Args[1]
	// grpc.WithTransportCredentials(insecure.NewCredentials())
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Set up a connection to the order server.
	conn, err := grpc.Dial(grpcServerEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}

	known_files := map[string]string{
		"/":                            "index.html",
		"/static/load-openapi.js":      "static/load-openapi.js",
		"/openapi/merged.swagger.json": "openapi/merged.swagger.json",
	}

	for path, file_path := range known_files {
		mux.HandlePath("GET", path, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			http.ServeFile(w, r, file_path)
		})
	}

	err = entities_api.RegisterServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = entities__0__reports__combined__xlsx.RegisterServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = entities__0__reports__summary.RegisterServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = reports__combined__zip.RegisterServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer conn.Close()
}
