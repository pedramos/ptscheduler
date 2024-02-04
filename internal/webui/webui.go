package webui

import (
	"embed"
	_ "embed"
	"net/http"

	"github.com/a-h/templ"
)

//go:generate tailwindcss --minify -o static/output.css
//go:generate templ generate

type NavBar struct{ Page, Link string }

type Pager interface {
	Header() templ.Component
	Body() templ.Component
}

func NewPage(nav []NavBar, p Pager) templ.Component {
	return Layout(nav, p.Header(), p.Body())
}

func NewPageHandler(nav []NavBar, p Pager) *templ.ComponentHandler {
	return templ.Handler(NewPage(nav, p))
}

//go:embed static
var staticfs embed.FS

func StaticHandler() http.Handler {
	return http.FileServer(http.FS(staticfs))
}
