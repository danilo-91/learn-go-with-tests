package concurrency

import (
	"reflect"
	"testing"
    "time"
)

func mockWebsiteChecker(url string) bool {
	if url == "https://mock.website.com" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://twitch.com",
		"https://mock.website.com",
	}

	got := CheckWebsite(mockWebsiteChecker, urls)
	want := map[string]bool{
		"https://google.com":       true,
		"https://twitch.com":       true,
		"https://mock.website.com": false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v but gotten %v", want, got)
	}
}

func slowWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func BenchmarkCheckWebsite(b *testing.B) {
    urls := make([]string, 100)
    for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        CheckWebsite(slowWebsiteChecker, urls)
    }
}
