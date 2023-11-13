package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		panic(err) // TODO: Remove the panic
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err) // TODO: remove the panic
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a href="mailto:tarunnahak25@gmail.com">Tarun Nahak</a></p>`)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<b>Q: Is there a free version?</b><br/>
	<p>A: Yes! we offer a free trial for 30 days on any paid plans.</p><br/>
	
	<b>Q: What are your support hours?</b><br/>
	<p>We have support staff answering emails 24/7, through response times may be a bit slower on weekends.</p><br/>
	
	<b>Q: How do I contact support</b><br/>
	<p>A: Email us <a emailto="support@lenslocked.com">support@lenslocked.com</a></p>`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) { http.Error(w, "Page not found", http.StatusNotFound) })
	fmt.Println("Starting the server on 3000...")
	http.ListenAndServe(":3000", r)
}
