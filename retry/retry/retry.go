// Package retry implements the Retry pattern using exponential backoff
// See https://cloud.google.com/iot/docs/how-tos/exponential-backoff
package retry

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
func Retry(f func() error) {
	var d time.Duration
	trial := uint8(0)

	err := f()
	for err != nil {
		fmt.Println(err)

		d = delay(trial)
		fmt.Printf("sleeping for %v...\n", d)
		time.Sleep(d)

		fmt.Println("retrying...")
		err = f()
		trial++
	}
}

func delay(n uint8) time.Duration {
	const maxBackoff = time.Second * (1 << 5)
	backoff := math.Min(math.Pow(2, float64(n)), float64(maxBackoff))
	return time.Duration(backoff)*time.Second + randMillis()
}

func randMillis() time.Duration {
	return time.Duration(rand.Intn(1000)) * time.Millisecond
}
