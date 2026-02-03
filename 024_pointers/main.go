package main

import "fmt"

func main() {
	// pointer is a something which stores memory address of any value

	//&x -> address of x( makes a pointer) (address operator)
	//*p -> go to the address and read/write (dereference operator)

	//when to use pointers
	//1. change the value in a separate function (by passing address) without returning it from the outer function
	// by default any variable is pass by value in golang(important) ( even you pass a array or slice that is also pass by value)

	score := 10

	fmt.Println("score:", score)
	addScore(&score)
	fmt.Println("score:", score)
}

func addScore(score *int) {
	*score += 10
}
