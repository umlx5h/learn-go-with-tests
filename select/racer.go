package main

import (
	"errors"
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	check := func(urls ...string) <-chan string {
		ret := make(chan string)

		for _, url := range urls {
			url := url
			go func() {
				defer close(ret)
				resp, err := http.Get(url)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == 200 {
					ret <- url
				}
			}()
		}

		return ret
	}

	result := check(url1, url2)
	firstReturned := <-result

	return firstReturned
}

func Racer2(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer3(a, b string, timeout time.Duration) (winner string, e error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("timeout exceeded")

	}
}

func ping(url string) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Head(url)
		close(ch)
	}()
	return ch
}

func main() {
}
