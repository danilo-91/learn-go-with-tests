package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	t        *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	t, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{t: t, mdParser: parser.New()}, nil
}

func (r *PostRenderer) Render(wr io.Writer, p Post) error {
	return r.t.ExecuteTemplate(wr, "post.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(wr io.Writer, p []Post) error {
	return r.t.ExecuteTemplate(wr, "index.gohtml", p)
}

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
