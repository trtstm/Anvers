package anvers

import "log"

type Plugin interface {
	Install(*Context)
	Uninstall(*Context)
	Activate(*Context)
	Deactivate(*Context)

	Name() string
}

type Registrar struct {
	plugins map[string]Plugin
	anvers  *Anvers
}

func NewRegistrar(a *Anvers) *Registrar {
	registrar := &Registrar{
		plugins: map[string]Plugin{},
		anvers:  a,
	}

	return registrar
}

func (r *Registrar) Load(plugins []Plugin) {
	for _, p := range plugins {
		if _, ok := r.plugins[p.Name()]; ok {
			panic("Plugin '" + p.Name() + "' already loaded.")
		}

		p.Install(&r.anvers.Context)
		log.Printf("Plugin '%s' installed.", p.Name())
		p.Activate(&r.anvers.Context)
		log.Printf("Plugin '%s' activated.", p.Name())
	}
}
