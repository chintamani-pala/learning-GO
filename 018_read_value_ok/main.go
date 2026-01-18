package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0,
	}

	fmt.Println("a", points["a"]) //10
	fmt.Println("b", points["b"]) //0, b is present as value 0
	fmt.Println("c", points["c"]) //0, c is not present as value 0  but still getting as zero which is confusing

	//to resolve this we use value, ok := points["c"]
	valueB, okB := points["b"]
	fmt.Println("b", valueB, okB) //0, true

	valueC, okC := points["c"]
	fmt.Println("c", valueC, okC) //0, false

	if value, ok := points["b"]; ok {
		fmt.Println("b is acutally present", value)
	}
	if value, ok := points["c"]; ok {
		fmt.Println("c is acutally present", value)
	} else {
		fmt.Println("c is not present")
	}

}
