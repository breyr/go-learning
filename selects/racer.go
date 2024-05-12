package selects

import (
	"fmt"
	"net/http"
	"time"
)

// := can only be using inside of a function
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// waiting on multiple channels, return whichever closes first
	// select blocks until then, both functions are called basically simultaneously
	// time.After also returns a chan
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// using a struct{} since there is no allocation, just need a signal that we are done
	// always have to make channels since the zero val for a channel is nil, blocking forever
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
