package main

import (
	"image"
	"image/draw"

	"github.com/faiface/gui"
)

func Display(env gui.Env) {

	loadPage := func(drw draw.Image) image.Rectangle {
		mainPage := image.Rect(0, 0, 900, 600)
		draw.Draw(drw, mainPage, image.White, image.ZP, draw.Src)

		return mainPage
	}

	env.Draw() <- loadPage

	for {
		select {
		case _, ok := <-env.Events():
			if !ok {
				close(env.Draw())
				return
			}

		}
	}
}