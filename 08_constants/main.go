package main

import "fmt"

func main() {
	const appName = "Go basics" //untyped string

	const maxUpload int = 25 //typed const integer

	const discountedPrice float64 = 10.3

	fmt.Println(appName, maxUpload, discountedPrice)
}
