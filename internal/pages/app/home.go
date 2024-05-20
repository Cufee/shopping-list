package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Home() (pages.Page, error) {
	return pages.NewPage(
		components.Container(
			h.H1(g.Text("App Home Page")),
			h.P(g.Text("This is a site showing off gomponents.")),
		),
	)
}
