package main

// import brings external packages in to the file that u are working on
//where actually you need it.
import (
	"fmt"
	"math"
)

func main() {
	//packageName.FunctionName -> call a function from a package
	fmt.Println("Hello World")
	fmt.Println(math.Sqrt(25))
}
