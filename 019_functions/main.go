package main

import "fmt"

func greet() {
	fmt.Println("Hello")
}

//single return function
func add(a int, b int) int {
	return a + b
}

//multiple return function
func sumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	greet()
	fmt.Println(add(1, 2))           //3
	fmt.Println(sumAndProduct(1, 2)) //3 2
	sum, product := sumAndProduct(1, 2)
	fmt.Println(sum, product) //3 2

	//ignoring return values
	sum, _ = sumAndProduct(3, 4)
	fmt.Println(sum) //7

	_, product = sumAndProduct(3, 4)
	fmt.Println(product) //12
}
