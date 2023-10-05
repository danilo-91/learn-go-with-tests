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

func Render(w io.Writer, p Post) error {
	t, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := t.Execute(w, p); err != nil {
		return err
	}

	return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}
