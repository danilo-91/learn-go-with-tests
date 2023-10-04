package blogrenderer

import (
    "io"
    "fmt"
)

func Render(w io.Writer, p Post) error {
    title := `<h1>%s</h1>
`
    desc := `<p>%s</p>
`
    tags := `<p>Tags: <ul>`

    _, err := fmt.Fprintf(w, title, p.Title)
    _, err = fmt.Fprintf(w, desc, p.Description)

    for _, t := range p.Tags {
        tags += fmt.Sprintf(`<li>%s</li>`, t)
    }
    tags += `</ul></p>`
    _, err = fmt.Fprintf(w, tags)

    return err
}

type Post struct {
    Title, Body, Description string
    Tags []string
}
