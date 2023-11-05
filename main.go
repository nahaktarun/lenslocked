package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site !</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a href="mailto:tarunnahak25@gmail.com">Tarun Nahak</a></p>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 		// OR
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprintf(w, "Page not found")
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		// OR
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "Page not found")
	}

}

func main() {
	var router Router
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/knowPath", pathHandler)
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", router)
}
