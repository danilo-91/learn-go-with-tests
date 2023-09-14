package main

import (
    "time"
    "net/http"
)

func Racer(url1, url2 string) string {
    d1 := getDuration(url1)
    d2 := getDuration(url2)

    if d1 < d2 {
        return url1
    }
    return url2
}

func getDuration(url string) time.Duration {
    start := time.Now()
    http.Get(url)
    return time.Since(start)
}
