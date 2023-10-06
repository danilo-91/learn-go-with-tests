package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"
	"testing/fstest"

	"github.com/approvals/go-approval-tests"
	"github.com/isedaniel/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		p = blogrenderer.Post{
			Title:       "Hello World",
			Description: "This is a description.",
			Tags:        []string{"go", "tdd"},
		}
	)

    const (
        post1Body = `This is the body of the post.`
    )
    fsys := fstest.MapFS{
        "post1.md": {Data: []byte(post1Body)},
    }

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("single post to HTML", func(t *testing.T) {
		b := bytes.Buffer{}
		if err := pr.Render(&b, p); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, b.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		p = blogrenderer.Post{
			Title:       "Hello, World!",
			Body:        "This is a post.",
			Description: "This is a description.",
			Tags:        []string{"go", "tdd"},
		}
	)
    pr, err := blogrenderer.NewPostRenderer()
    if err != nil {
        b.Fatal(err)
    }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pr.Render(io.Discard, p)
	}
}
