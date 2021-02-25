// Package retry implements the Retry pattern using exponential backoff
// See https://cloud.google.com/iot/docs/how-tos/exponential-backoff
package retry

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Retry exponentially retries the function until maxTrial or no errors
//
// Example output:
// > something went wrong
// > sleeping for 1.952s...
// > retrying...
// > something went wrong
// > sleeping for 2.86s...
// > retrying...
// > something went wrong
// > sleeping for 4.622s...
// > retrying...
// > it's ok
func Retry(maxTrial uint8, f func() error) error {
	trial := uint8(0)
	fmt.Println("Starting trials.")

	err := errors.New("dummy")
	for err != nil {
		if trial > maxTrial {
			return errors.New("trials exceeded")
		}
		err = f()
		if err == nil {
			fmt.Println("Finished.")
			return nil
		}
		fmt.Printf("err: %v\n", err)
		fmt.Println("Retrying...")

		d := delay(trial)
		fmt.Printf("Trial #%v, sleeping for %v...\n", trial, d)
		time.Sleep(d)

		trial++
	}
	return nil
}

func delay(n uint8) time.Duration {
	const maxBackoff = time.Second * (1 << 5)
	// limiting backoff
	backoff := math.Min(math.Pow(2, float64(n)), float64(maxBackoff))
	return time.Duration(backoff)*time.Second + randMillis()
}

func randMillis() time.Duration {
	// jitter, rand is for even distribution in range of a second
	return time.Duration(rand.Intn(1000)) * time.Millisecond
}
