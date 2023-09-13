package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  uint8
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// u := struct {
	// 	Name string
	// 	Age  uint8
	// 	Meta struct {
	// 		Visits int
	// 	}
	// }{
	// 	Name: "Emory.Du",
	// 	Age:  24,
	// 	Meta: struct{ Visits int }{
	// 		Visits: 4,
	// 	},
	// }
	u := User{
		Name: "Emory.Du",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!"); </script>`,
		Age:  24,
	}

	err = t.Execute(os.Stdout, u)
	if err != nil {
		panic(err)
	}
}
