package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestFail(t *testing.T) {
	startTime := time.Now()

	i := 0

	fn := func(attempt *Attempt) error {
		i = i + 1
		return fmt.Errorf("Error")
	}

	err := Retry(fn, 4, time.Millisecond*500)

	if err == nil {
		t.Error("Expected error")
	}

	if i != 4 {
		t.Error("Function did not run 4 times")
	}

	duration := time.Since(startTime)

	if duration.Seconds() < 3.5 {
		t.Error("Function should have slept for at least 3.5 seconds")
	}
}

func TestSuccess(t *testing.T) {
	fn := func(attempt *Attempt) error {
		return nil
	}

	err := Retry(fn, 4, time.Millisecond*500)

	if err != nil {
		t.Error(err)
	}
}

func TestRepeatOnFailure(t *testing.T) {
	i := 1

	fn := func(attempt *Attempt) error {
		if i != attempt.Attempt {
			t.Errorf("i (%d) did not match attempt (%d)", i, attempt.Attempt)
		}

		if i == 1 || i == 2 {
			i = i + 1
			return fmt.Errorf("An error")
		}

		return nil
	}

	err := Retry(fn, 5, time.Millisecond*500)

	if err != nil {
		t.Error("Unexpected error: %v", err)
	}

	if i != 3 {
		t.Error(fmt.Errorf("Function should have run 3 times, ran %d times", i))
	}
}

func TestAttemptStructure(t *testing.T) {
	fn := func(attempt *Attempt) error {
		if attempt.Attempt != 1 {
			t.Error("Should only run once")
		}

		if attempt.Total != 3 {
			t.Error("Total should be 3")
		}

		if attempt.Wait != 0 {
			t.Error("Initial wait should be 0")
		}

		return nil
	}

	Retry(fn, 3, time.Millisecond*500)

}
