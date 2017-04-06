package main

import (
	"github.com/trtstm/anvers"
	"github.com/trtstm/anvers/plugins/posts"
)

func main() {
	a := anvers.New(anvers.AnversConfig{
		Plugins: []anvers.Plugin{
			posts.NewPage(),
		},
	})

	a.Load()
}
