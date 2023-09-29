package main

import (
	"github.com/GodsBoss/pineapple-burger-pizza/pkg/game"

	"github.com/GodsBoss/gggg/v2/pkg/dom"
	"github.com/GodsBoss/gggg/v2/pkg/dominit"
)

func main() {
	win, _ := dom.GlobalWindow()
	doc, _ := win.Document()

	img, _ := doc.CreateImageElement("gfx.png")
	img.On(
		// Success
		func() {
			hint, _ := doc.GetElementByID("hint")
			dom.RemoveNode(hint)

			dominit.Run(game.New(img))
		},

		// Fail
		func(err interface{}) {},
	)

	<-make(chan struct{}, 0)
}
