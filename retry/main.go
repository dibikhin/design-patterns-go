// Package main contains an example of using the Retry pattern.
// Emulated errors are random.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	r "retry/retry"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func job() error {
	i := rand.Intn(2)
	if i == 0 {
		return errors.New("something went wrong, now: " + time.Now().String())
	}
	fmt.Println("It's ok.")
	return nil
}

// Run it multiple times to get the function retried due to its errors are random
func main() {
	maxTrial := uint8(1 << 5)
	err := r.Retry(maxTrial, job)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
