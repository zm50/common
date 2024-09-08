package tools

import (
	"errors"
	"time"
)

// Retry retries the given function tryCount times, and if an error occurs, it waits for 1 second before retrying.
func Retry(tryCount int, tryFunc func() error) error {
	return RetryWithHandleError(tryCount, tryFunc, func(error) {})
}

// RetryWithDelay retries the given function tryCount times, and if an error occurs, it waits for the given delay duration before retrying.
func RetryWithDelay(tryCount int, tryFunc func() error, delay time.Duration) error {
	return RetryWithHandleError(tryCount, tryFunc, func(error) { time.Sleep(delay) })
}

// RetryWithHandleError retries the given function tryCount times, and if an error occurs, it calls the handleError function with the error.
func RetryWithHandleError(tryCount int, tryFunc func() error, handleError func(error)) error {
	var err error
	for i := 0; i < tryCount; i++ {
		e := tryFunc()
		if e == nil {
			return nil
		}

		err = errors.Join(err, e)
	}

	return err
}
