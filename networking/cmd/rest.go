package main

import (
	"net/http"

	restHandler "user-service/internal/handler/rest"
	"user-service/internal/service"
)

func startREST(svc *service.UserService) *http.Server {
	mux := http.NewServeMux()
	restHandler.NewUserHandler(svc).Register(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go srv.ListenAndServe()

	return srv
}
