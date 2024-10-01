package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"reflect"

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
	// FIXME: there's probably a better way of checking whether it's an error message
	if reflect.TypeOf(v).String() == "*status.Status" {
		return "application/json"
	}

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

	registerFuncs := []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		entities_api.RegisterServiceHandler,
		entities__0__reports__combined__xlsx.RegisterServiceHandler,
		entities__0__reports__summary.RegisterServiceHandler,
		reports__combined__zip.RegisterServiceHandler,
	}

	for _, curr := range registerFuncs {
		err = curr(context.Background(), mux, conn)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}

	for path, data := range asset.Router {
		mux.HandlePath("GET", path, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			w.Write(data)
		})
	}

	serviceDescriptions := []grpc.ServiceDesc{
		entities_api.Service_ServiceDesc,
		entities__0__reports__combined__xlsx.Service_ServiceDesc,
		entities__0__reports__summary.Service_ServiceDesc,
		reports__combined__zip.Service_ServiceDesc,
	}

	var knownStreamingUrls = make(map[string]bool)

	for _, serviceDescription := range serviceDescriptions {
		for _, stream := range serviceDescription.Streams {
			urlPath := "/" + serviceDescription.ServiceName + "/" + stream.StreamName
			knownStreamingUrls[urlPath] = true
		}
	}

	supportNdjson := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// This has to be set so that the marshaller sees it
			if _, ok := knownStreamingUrls[r.URL.Path]; ok {
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
