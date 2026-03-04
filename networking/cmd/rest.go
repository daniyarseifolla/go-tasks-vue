package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	grpcHandler "user-service/internal/handler/grpc"
	"user-service/internal/service"
	"user-service/gen/pb"
)

func startREST(svc *service.UserService) *http.Server {
	mux := runtime.NewServeMux()
	pb.RegisterUserServiceHandlerServer(context.Background(), mux, grpcHandler.NewUserHandler(svc))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go srv.ListenAndServe()

	return srv
}
