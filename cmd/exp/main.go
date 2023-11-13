package main

import (
	"html/template"
	"os"
)

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// Ananymous struct
	user := struct {
		Name string
	}{Name: "Tarun"}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
