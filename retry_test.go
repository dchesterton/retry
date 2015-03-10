package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestFail(t *testing.T) {
	startTime := time.Now()

	fn := func(attempt *Attempt) error {
		return fmt.Errorf("Error")
	}

	err := Retry(fn, 4, time.Millisecond*500)

	if err == nil {
		t.Error("Expected error")
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

func TestAttemptStructure(t *testing.T) {
	fn := func(attempt *Attempt) error {
		if attempt.Attempt != 0 {
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
