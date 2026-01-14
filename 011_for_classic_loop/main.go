package main

import "fmt"

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println((i))
	}

	N := 10
	sum := 10
	for i := 1; i <= N; i++ {
		sum += i
	}
	fmt.Println("sum", sum)

}
