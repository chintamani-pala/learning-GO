package main

import "fmt"

func main() {
	//length and capacity
	//length -> number of elements
	//capacity -> how many number of elemnt you can store

	scores := make([]int, 0, 5)
	fmt.Println("Before appending ", scores, len(scores), cap(scores))
	scores = append(scores, 100)
	scores = append(scores, 2000, 3000, 4000, 5000, 60000, 60000, 60000, 60000, 60000, 60000, 60000, 60000)
	fmt.Println("After appending ", scores, len(scores), cap(scores))

}
