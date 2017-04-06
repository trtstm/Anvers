package posts

import "github.com/trtstm/anvers"

type Config struct {
	Name string
}

type Post struct {
	Config Config
}

func (p *Post) Install(ctx *anvers.Context) {
}

func (p *Post) Uninstall(ctx *anvers.Context) {
}

func (p *Post) Activate(ctx *anvers.Context) {
}

func (p *Post) Deactivate(ctx *anvers.Context) {
}

func (p *Post) Name() string {
	return p.Config.Name
}
