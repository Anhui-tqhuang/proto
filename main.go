package main

import (
	"context"
	"net"

	"github.com/TianqiuHuang/proto/pd/auth"
	"google.golang.org/grpc"
	"k8s.io/klog"
)

const (
	// PORT ...
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, &FakeAuthService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		klog.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}

// FakeAuthService ...
type FakeAuthService struct{}

// Validate ...
func (svc *FakeAuthService) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	// rawIdToken := req.GetRawIdToken()
	// validate rawIdToken
	claimNames := req.GetClaimNames()
	resp := &auth.ValidateResponse{
		Claims: map[string]*auth.Value{},
	}

	for _, claimName := range claimNames {
		var value = new(auth.Value)

		// switch data.(type) {
		// case string:
		// case int, int32, int64:
		// case []byte:
		// default:
		// }

		switch claimName {
		case "email":
			value = &auth.Value{
				Type: &auth.Value_Str{
					Str: "tianqiuhuang@gmail.com",
				},
			}
		}

		resp.Claims[claimName] = value
	}

	return resp, nil
}
