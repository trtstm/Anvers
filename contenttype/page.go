package contenttype

import "github.com/trtstm/anvers"

type Page struct {
	title string
	body  string
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
			{Title: "New Page", Link: "http://google.com"},
		},
	})
}

func (p *PagePlugin) Name() string {
	return "contenttype.page"
}

var pages = []Page{}
