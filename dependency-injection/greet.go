package main

import (
    "fmt"
    "log"
    "io"
    "net/http"
)

func Greet(w io.Writer, s string) {
    fmt.Fprintf(w, "Hello, %s", s)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, r.UserAgent())
}

func main() {
    log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
