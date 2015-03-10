package retry

import (
	"time"
)

type Attempt struct {
	Attempt int
	Total   int
	Wait    time.Duration
}
