package retry

import (
	"time"
)

type Attempt struct {
	Number   int
	NextWait time.Duration
}
