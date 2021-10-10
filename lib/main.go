package main

import (
	"github.com/akovardin/pong/components"
	"github.com/akovardin/pong/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

func main() {
	width := 20
	height := 20

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := ebiten.NewImageFromImage(image.NewRGBA(image.Rectangle{upLeft, lowRight}))

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, cyan)
		}
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(&game.Game{
		Xdirection: 1,
		Ydirection: 1,
		Width:      640,
		Height:     480,
		Gopher: &components.Gopher{
			Img:    img,
			X:      0,
			Y:      0,
			Width:  float64(width),
			Height: float64(height),
		},
	}); err != nil {
		log.Fatalln(err)
	}
}
