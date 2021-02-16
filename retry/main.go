// Package main contains an example of using the Retry pattern.
// Emulated errors are random.
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	r "retry/retry"
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
	r.Retry(something)
}
