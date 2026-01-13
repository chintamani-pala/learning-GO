package main

import "fmt"

func main() {
	var city string
	city = "Berhampur"

	var channelName = "ChintaCodeX" //infered to string

	//:= -> short way to declare and assign a variable

	followers := 100 //var followers int = 100

	followers = followers + 10

	likes, comments := 10, 20 //multiple variable declaration and assignment

	fmt.Println(city, channelName, followers, likes, comments)
}
