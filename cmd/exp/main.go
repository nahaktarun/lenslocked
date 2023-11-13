package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
	Bio  string
}

type UserMeta struct {
	Visits int
}

func main() {

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Tarun",
		Bio:  `<script>alert("hi there");</script>`,
		Age:  111,
		Meta: UserMeta{
			Visits: 34,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
