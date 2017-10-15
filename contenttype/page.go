package contenttype

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/trtstm/anvers"
)

type Page struct {
	Title string
	Body  string
}

// PagePlugin is a normal webpage.
type PagePlugin struct {
}

func (p *PagePlugin) Install(a *anvers.Anvers) {
}

func (p *PagePlugin) Register(a *anvers.Anvers) {
	a.AdminMenu.GetByName("content_types").AddEntry(&anvers.AdminMenuEntry{
		Title: "Page",
		Entries: []*anvers.AdminMenuEntry{
			{Title: "New Page", Link: "/admin/pages"},
		},
	})

	a.Router.Get("/admin/pages", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("../../assets/templates/admin/pages.html"))
		fmt.Println(t.Execute(w, map[string]interface{}{
			"Anvers": a,
			"pages":  pages,
		}))
	})
}

func (p *PagePlugin) Name() string {
	return "contenttype.page"
}

var pages = []*Page{
	&Page{Title: "Home page", Body: "<h1>Hello world</h1>"},
}
