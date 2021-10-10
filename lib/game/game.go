package game

import (
	"github.com/akovardin/pong/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const step = 10

type Game struct {
	Xdirection float64
	Ydirection float64
	Width  float64
	Height  float64
	Gopher  *components.Gopher
	Touches []ebiten.TouchID
}

func (g *Game) Update() error {
	// switch direction
	if g.Gopher.X+g.Gopher.Width > g.Width {
		g.Xdirection = -1
	} else if g.Gopher.X < 0 {
		g.Xdirection = 1
	}

	if g.Gopher.Y+g.Gopher.Height > g.Height {
		g.Ydirection = -1
	} else if g.Gopher.Y < 0 {
		g.Ydirection = 1
	}

	// input
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.Xdirection *= -1
		g.Ydirection *= -1
	}

	g.Touches = inpututil.AppendJustPressedTouchIDs(g.Touches)
	if len(g.Touches) > 0 {
		g.Xdirection *= -1
		g.Ydirection *= -1
		g.Touches = nil
	}

	// moving Gopher
	g.Gopher.X += g.Xdirection * step
	g.Gopher.Y += g.Ydirection * step

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Gopher.X, g.Gopher.Y)
	screen.DrawImage(g.Gopher.Img, op)
}

func (g *Game) Layout(w, h int) (int, int) {
	return int(g.Width), int(g.Height)
}
