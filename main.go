package main

import (
	"fmt"
	"html/template"
	"os"
	"reflect"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	tmpl, err := template.New("Tmpl1").Parse("Name: {{.Name}}\nEmail :{{.Email}}\nAge:{{.Age}}")
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------")
	fmt.Println("-----Parse----")
	user1 := &User{Name: "jhpark", Email: "jhpark@sinsiway.com", Age: 38}
	user2 := User{Name: "test2", Email: "test2@sinsiway.com", Age: 28}
	fmt.Println("type u1:", reflect.TypeOf(user1), " u2:", reflect.TypeOf(user2))
	tmpl.Execute(os.Stdout, user1)
	fmt.Println()
	tmpl.Execute(os.Stdout, user2)
	fmt.Println()

	fmt.Println("--------------")
	fmt.Println("--ParseFiles--")
	tmplFile, err := template.New("Tmpl2").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmplFile.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user1)
	if err != nil {
		panic(err)
	}
	err = tmplFile.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)
	if err != nil {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("-=---List---=-")
	user3 := User{Name: "jhpark", Email: "jhpark@sinsiway.com", Age: 38}
	users := []User{user2, user3}
	tmplListFile, err := template.New("Tmpl3").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl3.tmpl")
	if err != nil {
		panic(err)
	}
	err = tmplListFile.ExecuteTemplate(os.Stdout, "tmpl3.tmpl", users)
	if err != nil {
		panic(err)
	}
}
