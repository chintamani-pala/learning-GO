package main

import "fmt"

func main() {
	isLogged := true //infered as boolean type
	isAdmin := true
	hasSubscription := false

	//And operator (&&)
	canOpenDashboard := isLogged && hasSubscription
	//Or operator (||)
	canDeletePost := isAdmin || (isLogged && hasSubscription)
	fmt.Println(canOpenDashboard, canDeletePost)

	age := 25
	isAdult := age > 18
	fmt.Println(isAdult)
}
