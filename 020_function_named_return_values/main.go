package main

import "fmt"

func divide(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}
func main() {
	quotient, remainder := divide(10, 3)
	fmt.Println(quotient, remainder)
}
