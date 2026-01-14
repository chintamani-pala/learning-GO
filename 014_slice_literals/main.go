package main

import "fmt"

func main() {
	//common collection type
	//dynamic size
	// []type{...}

	results := []string{"chintamani", "pala"}
	fmt.Println(results[len(results)-1])

	var nums []int

	nums = append(nums, 1)
	nums = append(nums, 2, 3, 4, 5)
	nums = append(nums, 3)

	fmt.Println(nums)
}
