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

	err := template.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func signInGetHandler(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("./templates/login.html"))

	err := template.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func signInPostHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form")
	}

	fmt.Printf("Email: %s \n", r.Form["email"][0])
	fmt.Printf("Password: %s \n", r.Form["password"][0])

	template := template.Must(template.ParseFiles("./templates/index.html"))
	template.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func signOutGetHandler(w http.ResponseWriter, r *http.Request) {

}

func signOutPostHandler(w http.ResponseWriter, r *http.Request) {

}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("./templates/register.html"))

	err := template.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
	r.Get("/login", signInGetHandler)
	r.Post("/login", signInPostHandler)
	r.Get("/register", registerGetHandler)

	log.Fatal(http.ListenAndServe(":3060", r))
}
