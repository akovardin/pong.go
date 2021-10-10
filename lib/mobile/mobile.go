package mobile

import (
	"github.com/akovardin/pong/components"
	"github.com/akovardin/pong/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"image"
	"image/color"
)

func init() {
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

	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	mobile.SetGame(&game.Game{
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
	})
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
