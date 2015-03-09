package retry

import (
	"math"
	"time"
)

type retryFunction func(attempt *Attempt) error

func Retry(fn retryFunction, tries int, startWait time.Duration) error {
	var err error
	attemptNumber := 0

	for {
		wait := startWait * time.Duration(math.Pow(2, float64(attemptNumber)))

		attempt := &Attempt{
			Number:   attemptNumber,
			NextWait: wait,
		}

		err = fn(attempt)

		if err == nil {
			return nil
		}

		attemptNumber += 1

		if attemptNumber == tries {
			return err
		}

		time.Sleep(wait)
	}
}
