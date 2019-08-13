package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/grpc"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func validID(id int) bool {
	valid := []int{1,160401}
	return contains(valid, id)
}

func fetchToken(id int) string {
	return "some test token"
}

// empty struct because this isn't a fancy example
type AuthorizationServer struct{}

// Inject header with API token if ID is valid
func (authserv *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string
	splitToken = strings.Split(authHeader, "Bearer ")
	if ok && len(splitToken) == 2 {
		// Extract Social Network ID from auth
		sn_id, _ := strconv.Atoi(splitToken[1])

		// This is where you'd go check with the system that knows if it's a valid token.
		if validID(sn_id) {
			tokenSha := fetchToken(sn_id)
			return &auth.CheckResponse{
				Status: &rpc.Status{
					Code: int32(rpc.OK),
				},
				HttpResponse: &auth.CheckResponse_OkResponse{
					OkResponse: &auth.OkHttpResponse{
						Headers: []*core.HeaderValueOption{
							{
								Header: &core.HeaderValue{
									Key:   "x-ext-auth-token",
									Value: tokenSha,
								},
							},
						},
					},
				},
			}, nil
		}
	}

	return &auth.CheckResponse{
		Status: &rpc.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
				Body: "Invalid Auth Header!",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	grpcServer := grpc.NewServer()
	authServer := &AuthorizationServer{}
	auth.RegisterAuthorizationServer(grpcServer, authServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
