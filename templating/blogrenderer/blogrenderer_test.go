package blogrenderer_test

import (
    "bytes"
    "testing"
    "github.com/isedaniel/blogrenderer"
)

func TestRender(t *testing.T) {
    var (
        aPost = blogrenderer.Post{
            Title: "hello world",
            Body: "This is a post.",
            Description: "This is a description",
            Tags: []string{"go", "tdd"},
        }
    )

    t.Run("single post to HTML", func(t *testing.T) {
        b := bytes.Buffer{}
        err := blogrenderer.Render(&b, aPost)

        if err != nil {
            t.Fatal(err)
        }

        got := b.String()
        want := `<h1>hello world</h1>`

        if got != want {
            t.Errorf("wanted %q, but got %q", want, got)
        }
    })
}
