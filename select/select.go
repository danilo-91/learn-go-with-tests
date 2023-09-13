package main

import (
    "time"
    "net/http"
)

func Racer(url1, url2 string) string {
    startURL1 := time.Now()
    http.Get(url1)
    duration1 := time.Since(startURL1)

    startURL2 := time.Now()
    http.Get(url2)
    duration2 := time.Since(startURL2)

    if duration1 < duration2 {
        return url1
    }
    return url2
}
