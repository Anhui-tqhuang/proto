package main

import (
	"context"
	"fmt"

	"github.com/TianqiuHuang/proto/pd/auth"
	"google.golang.org/grpc"
	"k8s.io/klog"
)

const (
	// PORT ...
	PORT = "9002"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		klog.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)
	resp, err := client.Validate(context.Background(), &auth.ValidateRequest{
		ClaimNames: []string{"email"},
	})
	if err != nil {
		klog.Fatalf("grpc.Dial err: %v", err)
	}

	fmt.Printf("email: '%s'\n", resp.Claims["email"].GetStr())
}
