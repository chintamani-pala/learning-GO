package main

import "fmt"

func main() {
	//map[keyType]valueType
	ages := map[string]int{
		"Chintamani": 22,
		"Gobinda":    25,
		"Asutosh":    22,
	}

	fmt.Println(ages)               //map[Asutosh:22 Chintamani:22 Gobinda:25]
	fmt.Println(len(ages))          //3
	fmt.Println(ages["chintamani"]) // `0` because of key is not same
	fmt.Println(ages["Chintamani"]) // `22` because of key is not same

	//using make(map[keyType]valueType)

	var scores map[string]int
	fmt.Println(scores)      //map[]
	fmt.Println(len(scores)) //0

	scores = make(map[string]int)
	scores["math"] = 90
	scores["english"] = 80
	scores["science"] = 70
	fmt.Println(scores)         //map[Chintamani:22]
	fmt.Println(len(scores))    //1
	fmt.Println(scores["math"]) //90

	users := map[string]string{
		"u1": "sangam",
		"u2": "chintamani",
		"u3": "gobinda",
	}
	fmt.Println(users) //map[u1:sangam u2:chintamani u3:gobinda]

	//delete a key from map
	delete(users, "u2")
	fmt.Println(users) //map[u1:sangam u3:gobinda]

	delete(users, "nokey") //no error although key is not present
	fmt.Println(users)     //map[u1:sangam u3:gobinda]

	// for range in map
	prices := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 30,
	}

	total := 0
	for item, price := range prices {
		total += price
		fmt.Println(item, price)
	}
	fmt.Println(total)

	// only keys from map
	for item := range prices {
		fmt.Println(item)
	}

	//only value from
	for _, price := range prices {
		fmt.Println(price)
	}

}
