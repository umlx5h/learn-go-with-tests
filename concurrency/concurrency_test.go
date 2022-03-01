package concurrency

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsitesConc2(mockWebsiteChecker, websites)

	assert.Equal(t, want, got)
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		// urls[i] = "a url"
		urls[i] = fmt.Sprintf("a url %v", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsitesConc2(slowStubWebsiteChecker, urls)
	}
}
