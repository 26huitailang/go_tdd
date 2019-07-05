package racer

import (
	"net/http"
	"time"
)

func Racer(u1, u2 string) (winner string) {
	aDuration := measureResponseTime(u1)
	bDuration := measureResponseTime(u2)

	if aDuration < bDuration {
		return u1
	}
	return u2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}
