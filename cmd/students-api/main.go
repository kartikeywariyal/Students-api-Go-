package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kartikeywariyal/students-api-Go-/internal/config"
	"github.com/kartikeywariyal/students-api-Go-/internal/http/handlers/student"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	router.HandleFunc("GET /job", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is ur first JOb")
	})
	server := &http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Println(cfg.HttpServer.Address)
	fmt.Println("Starting server on", cfg.HttpServer.Address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		server.ListenAndServe()
	}()
	<-done
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		panic("server failed to shut down: " + err.Error())
	}
	fmt.Println("Server stopped")

}
