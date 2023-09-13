package concurrency

type WebsiteChecker func(string) bool

type result struct {
    string
    bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
    resultCh := make(chan result)

	for _, url := range urls {
        go func(s string){
            resultCh <- result{s, wc(s)}
        }(url)
	}

    for range urls {
        r := <-resultCh
        results[r.string] = r.bool
    }

	return results
}
