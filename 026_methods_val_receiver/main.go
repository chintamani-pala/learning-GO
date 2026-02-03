package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Chintamani",
		Age:  22,
	}
	fmt.Println(user.Intro())
}

//value receiver means this method receives a copy of the user
//so it cannot modify the original user
func (u User) Intro() string {
	return fmt.Sprintf("Hello, My name is %s and I am %d years old", u.Name, u.Age)
}
