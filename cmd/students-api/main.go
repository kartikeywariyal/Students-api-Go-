package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kartikeywariyal/students-api-Go-/internal/config"
	"github.com/kartikeywariyal/students-api-Go-/internal/http/handlers/student"
	"github.com/kartikeywariyal/students-api-Go-/internal/storage/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.MustLoad()
	storage, err := sqlite.NewSqliteStorage(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetStudent(storage))
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
	err = server.Shutdown(ctx)
	if err != nil {
		panic("server failed to shut down: " + err.Error())
	}
	fmt.Println("Server stopped")

}
