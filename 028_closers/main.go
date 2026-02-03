package main

import "fmt"

func counter() func() int {
	count := 0

	//return a anonymous function which has access to count variable
	//even after the counter function returns
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("First counter")
	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println("First counter End")

	fmt.Println("Second counter")
	counter2 := counter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println("Second counter End")
}
