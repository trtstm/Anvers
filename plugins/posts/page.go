package posts

import (
	"fmt"

	"github.com/trtstm/anvers"
)

type page struct {
	Post
}

func NewPage() *page {
	p := &page{
		Post: Post{Config: Config{
			Name: "anvers-page",
		}},
	}

	return p
}

func (p *page) Install(ctx *anvers.Context) {
	fmt.Println("page")
	p.Post.Install(ctx)
}
