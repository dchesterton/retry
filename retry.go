package retry

import (
	"math"
	"time"
)

type retryFunction func(attempt *Attempt) error

func Retry(fn retryFunction, attempts int, startWait time.Duration) error {
	var err error
	var wait time.Duration

	for i := 1; i <= attempts; i++ {
		if i == 1 {
			wait = 0
		} else {
			wait = startWait * time.Duration(math.Pow(2, float64(i-2)))
			time.Sleep(wait)
		}

		err = fn(&Attempt{
			Attempt: i,
			Total:   attempts,
			Wait:    wait,
		})

		if err == nil {
			return nil
		}
	}

	return err
}
