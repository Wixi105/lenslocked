package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err.Error())
	}
	
	user := User{Name: "Ronald McDonald"}
	
	// Anonymous struct
	// user2:= struct {
	// 	Name string
	// }{
	// 	Name: "Susan Smith",
	// }

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err.Error())
	}
}
