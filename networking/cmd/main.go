package main

import (
	"context"
	"os/signal"
	"syscall"

	"user-service/internal/repository"
	"user-service/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)

	grpcSrv := startGRPC(svc)
	httpSrv := startREST(svc)

	<-ctx.Done()

	grpcSrv.GracefulStop()
	httpSrv.Shutdown(context.Background())
}
