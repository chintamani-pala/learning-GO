package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Chintamani"
	lastName := "Pala"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}
