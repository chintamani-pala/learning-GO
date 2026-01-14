package main

import "fmt"

func main() {
	views := 1000
	views2 := 1000
	totalViews := views + views2

	likes := 10
	likes++
	likes += 10

	avgViews := totalViews / 2

	fmt.Println(totalViews, avgViews, likes)

	rating1 := 4.5
	rating2 := 5.1

	avgRating := rating1 + rating2

	fmt.Println(avgRating)
}
