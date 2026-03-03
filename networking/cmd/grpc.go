package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	grpcHandler "user-service/internal/handler/grpc"
	"user-service/internal/service"
	"user-service/pkg/pb"
)

func startGRPC(svc *service.UserService) *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterUserServiceServer(srv, grpcHandler.NewUserHandler(svc))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on :50051: %v", err)
	}

	go srv.Serve(lis)

	return srv
}
