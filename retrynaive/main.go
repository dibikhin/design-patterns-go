// Package main contains a naive implementation of the Retry pattern.
// No limits, hardcoded delay, linear time.
// Emulated errors are random.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func something() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	if i == 0 {
		return errors.New("something went wrong")
	}
	fmt.Println("it's ok")
	return nil
}

// Run it multiple times to get the function retried due to its errors are random
func main() {
	err := something()
	for err != nil {
		fmt.Println(err)

		delay := time.Second
		fmt.Printf("sleeping for %v...\n", delay)
		time.Sleep(delay)
		fmt.Println("retrying...")

		err = something()
	}
}
