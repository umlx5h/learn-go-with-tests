package concurrency

import (
	"sync"
)

type WebsiteChecker func(string) bool

type Result struct {
	url    string
	result bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

func CheckWebsitesConc(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	var wg sync.WaitGroup
	wg.Add(len(urls))
	resultCh := make(chan Result, len(urls))
	for _, url := range urls {
		url := url
		go func() {
			defer wg.Done()
			r := wc(url)

			resultCh <- Result{url: url, result: r}
		}()
	}

	wg.Wait()
	close(resultCh)

	for r := range resultCh {
		results[r.url] = r.result
	}

	return results
}

func CheckWebsitesConc2(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	type result struct {
		string
		bool
	}

	resultChannel := make(chan result)

	for _, url := range urls {
		url := url
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

func mWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	CheckWebsitesConc(mWebsiteChecker, websites)
}
