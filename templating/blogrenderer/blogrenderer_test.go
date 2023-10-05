package blogrenderer_test

import (
	"bytes"
	"github.com/isedaniel/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		p = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post.",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("single post to HTML", func(t *testing.T) {
		b := bytes.Buffer{}
		err := blogrenderer.Render(&b, p)

		if err != nil {
			t.Fatal(err)
		}

		got := b.String()
		want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("wanted %q, but got %q", want, got)
		}
	})
}
