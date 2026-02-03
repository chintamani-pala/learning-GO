package main

import "fmt"

//struct groups related fields into one type
// In struct fields are by default non required -> to achive required we use constructor functions( IMportant learn this ASAP)
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {

	//create struct value
	user1 := User{
		ID:    1,
		Name:  "John",
		Email: "john@mail.com",
		Age:   30,
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)
	fmt.Println(user1.Age)

	//update Age of user (struct is mutatble by default)
	user1.Age = 31
	fmt.Println(user1.Age)

	//partial user
	user2 := User{
		Name: "Jane",
	}
	fmt.Println(user2)

	user2.Email = "[EMAIL_ADDRESS]"
	fmt.Println(user2.Email)

	//anonymous struct
	user3 := struct {
		Name string
		Age  int
	}{
		Name: "Jane",
		Age:  25,
	}
	fmt.Println(user3)

}
