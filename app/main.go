package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&pageView{})
}

type pageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *pageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello Vecty!"),
	)
}
