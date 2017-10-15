package main

import (
	"github.com/trtstm/anvers/contenttype"

	"github.com/trtstm/anvers"
)

func main() {
	cms := anvers.NewAnvers(&anvers.Config{
		Plugins: []anvers.Plugin{
			&contenttype.PagePlugin{},
		},
	})

	cms.Start()
}
