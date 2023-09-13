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

	u := struct {
		Name string
		Age  uint8
		Meta struct {
			Visits int
		}
	}{
		Name: "Emory.Du",
		Age:  24,
		Meta: struct{ Visits int }{
			Visits: 4,
		},
	}

	err = t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
