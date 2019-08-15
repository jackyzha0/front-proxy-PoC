package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	rpc "istio.io/gogo-genproto/googleapis/google/rpc"
)

type AuthorizationServer struct{
	socialProfileService SocialProfileService
}

func constructgRPCResponse(status int32, success bool, envoyStatus int, body string, headers []*core.HeaderValueOption) (*auth.CheckResponse, error) {
	if success {
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: status,
			},
			HttpResponse: &auth.CheckResponse_OkResponse{
				OkResponse: &auth.OkHttpResponse{
					Headers: headers,
				},
			},
		}, nil
	} else {
		return &auth.CheckResponse{
			Status: &rpc.Status{
				Code: status,
			},
			HttpResponse: &auth.CheckResponse_DeniedResponse{
				DeniedResponse: &auth.DeniedHttpResponse{
					Status: &envoy_type.HttpStatus{
						Code: envoy_type.StatusCode(int32(envoyStatus)),
					},
					Body: body,
				},
			},
		}, nil
	}
}

// Inject header with API token if ID is valid
func (authserv *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string
	splitToken = strings.Split(authHeader, "Bearer ")
	if ok && len(splitToken) == 2 {
		// Extract Social Network ID from auth
		sn_id, _ := strconv.Atoi(splitToken[1])

		profile, err := authserv.socialProfileService.GetSocialProfile(int64(sn_id))

		if err != nil {
			return constructgRPCResponse(int32(rpc.INVALID_ARGUMENT), false, http.StatusFailedDependency, "Invalid Auth Header!", nil)
		}

		headers := []*core.HeaderValueOption{
			{
				Header: &core.HeaderValue{
					Key:   "x-ext-auth1",
					Value: profile.Auth1,
				},
			}, {
				Header: &core.HeaderValue{
					Key:   "x-ext-auth2",
					Value: profile.Auth2,
				},
			}, {
				Header: &core.HeaderValue{
					Key:   "x-ext-sn-auth1",
					Value: profile.SocialNetworkAppAuth1,
				},
			}, {
				Header: &core.HeaderValue{
					Key:   "x-ext-sn-auth2",
					Value: profile.SocialNetworkAppAuth2,
				},
			}, {
				Header: &core.HeaderValue{
					Key:   "x-ext-sn-auth3",
					Value: profile.SocialNetworkAppAuth3,
				},
			}, {
				Header: &core.HeaderValue{
					Key:   "x-ext-sn-app-id",
					Value: strconv.FormatUint(profile.SocialNetworkAppID, 10),
				},
			},
		}
		return constructgRPCResponse(int32(rpc.OK), true, -1, "", headers)
	}
	return constructgRPCResponse(int32(rpc.UNAUTHENTICATED), false, 401, "Invalid Auth Header!", nil)
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	grpcServer := grpc.NewServer()
	authServer := &AuthorizationServer{NewSocialProfileService()}
	auth.RegisterAuthorizationServer(grpcServer, authServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
