package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func hompageHandler(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("./templates/index.html"))

	err := template.Execute(w, "home")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func signInGetHandler(w http.ResponseWriter, r *http.Request) {

}

func signInPostHandler(w http.ResponseWriter, r *http.Request) {

}

func signOutGetHandler(w http.ResponseWriter, r *http.Request) {

}

func signOutPostHandler(w http.ResponseWriter, r *http.Request) {

}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {

}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {

}

func privatePage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Session auth")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", hompageHandler)

	log.Fatal(http.ListenAndServe(":3060", r))
}
