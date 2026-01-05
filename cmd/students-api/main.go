package main

import (
	"fmt"
	"net/http"

	"github.com/kartikeywariyal/students-api-Go-/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET  /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Students API")
	})
	router.HandleFunc("GET /job", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is ur first JOb")
	})
	server := &http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Println(cfg.HttpServer.Address)
	fmt.Println("Starting server on", cfg.HttpServer.Address)
	err := server.ListenAndServe()
	if err != nil {
		panic("server failed to start: " + err.Error())
	}

}
