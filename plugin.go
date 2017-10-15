package anvers

// Plugin is the generic interface for plugins.
type Plugin interface {
	Install(*Anvers)
	Register(*Anvers)
}
