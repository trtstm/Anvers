package anvers

type AnversConfig struct {
	Plugins []Plugin
}

type Context struct {
	Anvers          *Anvers
	PluginRegistrar *Registrar
}

type Anvers struct {
	pluginRegistrar *Registrar
	config          AnversConfig

	Context Context
}

func New(config AnversConfig) *Anvers {
	a := &Anvers{}
	a.pluginRegistrar = NewRegistrar(a)
	a.config = config
	a.Context = Context{
		Anvers:          a,
		PluginRegistrar: a.pluginRegistrar,
	}

	return a
}

func (a *Anvers) PluginRegistrar() *Registrar {
	return a.pluginRegistrar
}

func (a *Anvers) Load() {
	a.pluginRegistrar.Load(a.config.Plugins)
}
