package blogrenderer

import (
    "io"
    "fmt"
)

func Render(w io.Writer, p Post) error {
    _, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
    return err
}

type Post struct {
    Title, Body, Description string
    Tags []string
}
