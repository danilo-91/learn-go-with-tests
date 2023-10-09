package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

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

	if err := r.t.ExecuteTemplate(wr, "post.gohtml", p); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(wr io.Writer, p []Post) error {
    indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`
    t, err := template.New("index").Parse(indexTemplate)
    if err != nil {
        return err
    }

    if err := t.Execute(wr, p); err != nil {
        return err
    }
    return nil
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
    return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
