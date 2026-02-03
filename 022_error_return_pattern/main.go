package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// go dont use exceptions for normal use
	// functions -> return errors as normal return values

	// val, err := something()
	// if err != nil {
	// 	handle error
	// }

	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {

	input := "6"

	level, err := parseLevel(input)
	if err != nil {
		return err
	}

	fmt.Println("level:", level)
	return nil

}

func parseLevel(s string) (int, error) {
	// (value, error) -> return pattrn
	// nil error -> success
	// non-nil error -> failure

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid level: %s", s)
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level out of  (1-5)range: %d", n)
	}
	return n, nil
}
