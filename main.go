package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func hompageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homepage reached"))
}

func main() {
	fmt.Println("Session auth")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", hompageHandler)

	log.Fatal(http.ListenAndServe(":3060", r))
}
