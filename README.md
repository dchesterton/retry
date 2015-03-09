# Go retry
A simple Go function for retrying functions on error, with exponential backoff.

Example
===========

    fn := func(attempt *Attempt) error {
        return functionWhichMightError()
    }

    err := Retry(fn, 5, time.Millisecond*500)

The code will run up to 5 times, until `err == nil`, with a delay between each run. The delay will increase exponentially. In this example, the delays will be `500ms`, `1s`, `2s`, `4s`.
