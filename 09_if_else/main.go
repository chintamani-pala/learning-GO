package main

import "fmt"

func main() {
	score := 91

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 75 {
		fmt.Println("Grade B")
	} else if score >= 45 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Fail")
	}
}
