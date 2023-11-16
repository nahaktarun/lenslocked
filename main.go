package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	controller "github.com/nahaktarun/lenslocked/controllers"
	"github.com/nahaktarun/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controller.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controller.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controller.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
