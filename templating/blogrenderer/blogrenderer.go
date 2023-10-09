package blogrenderer

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/isedaniel/md"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	t *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	t, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{t: t}, nil
}

func (r *PostRenderer) Render(wr io.Writer, p Post) error {
    // Render body from MD to HTML
    newBody := string(md.MdToHtml([]byte(p.Body)))
    p.Body = newBody
    fmt.Println(p.Body)

	if err := r.t.ExecuteTemplate(wr, "post.gohtml", p); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}
