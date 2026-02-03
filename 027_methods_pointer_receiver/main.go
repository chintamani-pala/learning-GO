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
	fmt.Println("Before : ", user.Age)
	user.BirthDay()
	fmt.Println("After : ", user.Age)
}

// pointer receiver means this method receives a pointer to the user
// so it can modify the original user
func (u *User) BirthDay() {
	u.Age++
}
