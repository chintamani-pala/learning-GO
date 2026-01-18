package main

import "fmt"

func main() {
	views := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for range over slices
	total := 0
	for index, value := range views {
		total += value
		fmt.Println(index, value)
	}
	fmt.Println(total)
}
