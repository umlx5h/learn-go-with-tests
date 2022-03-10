package blogrenderer

import (
	"embed"
	"io"
	"text/template"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	t, err := template.ParseFS(postTemplates, "templates/*.go.tpl")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: t}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	// if err := r.templ.Execute(w, p); err != nil {
	// 	return err
	// }

	// return nil

	return r.templ.ExecuteTemplate(w, "blog.go.tpl", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	// indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`

	// templ, err := template.New("index").Parse(indexTemplate)
	// templ, err := template.New("index").Funcs(template.FuncMap{
	// 	"sanitiseTitle": func(title string) string {
	// 		return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	// 	},
	// }).Parse(indexTemplate)
	// if err != nil {
	// 	return err
	// }

	// if err := templ.Execute(w, posts); err != nil {
	// 	return err
	// }

	// return nil

	return r.templ.ExecuteTemplate(w, "index.go.tpl", posts)
}
