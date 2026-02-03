package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Sucess work")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Failed work")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {
	//resource related
	//start message -> resource acruired
	//work
	//end message -> resource released

	fmt.Println("starting work")

	//defer guarenties to run before function exit
	//even if function returns or panics

	defer fmt.Println("Cleanup: releasing resource")

	if !success {
		fmt.Println("work failed")
		return errors.New("work failed")
	}

	fmt.Println("work completed")
	return nil
}
