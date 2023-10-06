package blogrenderer

import (
	"embed"
	"html/template"
	"io"
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
	if err := r.t.ExecuteTemplate(wr, "post.gohtml", p); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}
