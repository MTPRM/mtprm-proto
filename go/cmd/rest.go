package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/mtprm/mtprm-proto-grpc-gateway/asset"
	entities_api "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities/v1"
	entities__0__reports__combined__xlsx "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities__0__reports__combined__xlsx/v1"
	entities__0__reports__summary "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/entities__0__reports__summary/v1"
	reports__combined__zip "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/beta/resources/reports__combined__zip/v1"
)

const ndJsonMimeType = "application/x-ndjson"

type NdJsonMarshaller struct {
	runtime.JSONPb
}

func (m *NdJsonMarshaller) ContentType(v interface{}) string {
	return ndJsonMimeType
}

func main() {
	mux := runtime.NewServeMux(
		// This only gets called if `Accept` header matches
		runtime.WithMarshalerOption(ndJsonMimeType, &NdJsonMarshaller{}),
	)

	grpcServerEndpoint := os.Args[1]
	// grpc.WithTransportCredentials(insecure.NewCredentials())
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Set up a connection to the order server.
	conn, err := grpc.Dial(grpcServerEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
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

	for path, data := range asset.Router {
		mux.HandlePath("GET", path, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.Write(data)
		})
	}

	supportNdjson := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// This has to be set so that the marshaller sees it
			if strings.HasSuffix(r.URL.Path, ".Service/List") {
				r.Header.Set("Accept", ndJsonMimeType)
			}

			h.ServeHTTP(w, r)
		})
	}

	port := ":8080"

	err = http.ListenAndServe(port, supportNdjson(mux))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Listening on port:", port)

	defer conn.Close()
}
