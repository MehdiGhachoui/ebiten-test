package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	img  *ebiten.Image
	posR = 50
	posL = 50
)

func init() {
	var err error

	img, err = ebitenutil.NewImageFromURL("https://ebitengine.org/go/renderimage/gopher.png")

	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

// Movement based on frames [inpututil]
func playerMovement(key ebiten.Key) bool {
	const (
		delay    = 1
		interval = 5
	)

	//Check the key press in the current frame
	d := inpututil.KeyPressDuration(key)

	//Pressed 1 time
	if d == 1 {
		return true
	}

	// Wait for "delay" then repeat after every interval "frame"
	if d >= delay && d%interval == 0 {
		return true
	}

	return false
}

// Always run on 60fps regardless of machine if not then it's an engine problem
func (g *Game) Update() error {
	//Smooth Movement
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		posR += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		posR -= 10
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posR), float64(posL))
	op.GeoM.Scale(1.5, 1)
	screen.DrawImage(img, op)
}

// works like zoom somehow ??
// more meangninful when the window is resizable
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
