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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<b>Q: Is there a free version?</b><br/>
	<p>A: Yes! we offer a free trial for 30 days on any paid plans.</p><br/>
	
	<b>Q: What are your support hours?</b><br/>
	<p>We have support staff answering emails 24/7, through response times may be a bit slower on weekends.</p><br/>
	
	<b>Q: How do I contact support</b><br/>
	<p>A: Email us <a emailto="support@lenslocked.com">support@lenslocked.com</a></p>`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
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
