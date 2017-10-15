package anvers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

// Config contains configuration options for the cms.
type Config struct {
	Plugins []Plugin
}

// Anvers is our main struct. It wraps the cms.
type Anvers struct {
	plugins   []Plugin
	AdminMenu *AdminMenuEntry
	Router    *chi.Mux
}

// NewAnvers creates a new cms.
func NewAnvers(config *Config) *Anvers {
	a := &Anvers{
		plugins:   []Plugin{},
		AdminMenu: newAdminMenu(),
		Router:    chi.NewRouter(),
	}

	for _, p := range config.Plugins {
		p.Install(a)
		a.plugins = append(a.plugins, p)
	}

	return a
}

// Start starts the cms.
func (a *Anvers) Start() {
	for _, p := range a.plugins {
		p.Register(a)
	}

	a.Router.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("../../assets/templates/admin/dashboard.html"))
		fmt.Println(t.Execute(w, map[string]interface{}{
			"Anvers": a,
		}))
	})

	http.ListenAndServe(":3000", a.Router)
}
