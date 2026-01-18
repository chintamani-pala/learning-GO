package main

import "fmt"

func sumAll(numbers ...int) int { // ... is variadic parameter
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	values := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sumAll(values...)) //expand the slice into individual values

	res := func(n int) int { //this is defining an anonymous function and assigned to a variables
		return n * n
	}
	fmt.Println(res(2))

	//IIFE in golang (Immediately Invoked Function Expression)
	res1 := func(a int, b int) int {
		return a + b
	}(2, 4)
	fmt.Println(res1)
}
