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
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/mtprm/mtprm-proto-grpc-gateway/asset"

	api_entities_v1 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entities/v1"
	api_entities__0__reports__combined__xlsx_v1 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entities__0__reports__combined__xlsx/v1"
	api_entities__0__reports__summary_v1 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entities__0__reports__summary/v1"
	api_entity_inquiries_v1 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entity_inquiries/v1"
	api_entity_inquiries_v2 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entity_inquiries/v2"
	api_entity_inquiries_v3 "github.com/mtprm/mtprm-proto-grpc-gateway/generated/mtprm/api/portfolio/resources/entity_inquiries/v3"
)

const ndJsonMimeType = "application/x-ndjson"

type NdJsonMarshaller struct {
	runtime.JSONPb
}

func (m *NdJsonMarshaller) ContentType(v interface{}) string {
	// status code: 401
	//
	// ```
	// {
	//   "error": {
	//     "code": 16,
	//     "message": "UnauthenticatedException: Missing or invalid authentication"
	//   }
	// }
	// ```
	if _, ok := v.(map[string]protoreflect.ProtoMessage); ok {
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
		api_entities_v1.RegisterServiceHandler,
		api_entities__0__reports__combined__xlsx_v1.RegisterServiceHandler,
		api_entities__0__reports__summary_v1.RegisterServiceHandler,
		api_entity_inquiries_v1.RegisterServiceHandler,
		api_entity_inquiries_v2.RegisterServiceHandler,
		api_entity_inquiries_v3.RegisterServiceHandler,
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
		api_entities_v1.Service_ServiceDesc,
		api_entities__0__reports__combined__xlsx_v1.Service_ServiceDesc,
		api_entities__0__reports__summary_v1.Service_ServiceDesc,
		api_entity_inquiries_v1.Service_ServiceDesc,
		api_entity_inquiries_v2.Service_ServiceDesc,
		api_entity_inquiries_v3.Service_ServiceDesc,
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
