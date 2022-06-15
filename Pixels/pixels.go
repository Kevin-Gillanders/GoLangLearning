package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var black color.Color
var red   color.Color
var blue  color.Color
var green color.Color
var rect Rect

func init(){
	red   = color.RGBA{255,221,153, 255}
	black = color.RGBA{0, 0, 0, 255}
	blue  = color.RGBA{0, 128, 128, 255}
	green = color.RGBA{255,128,102, 255}
	rect = newRect(10, 10, 25, 25)
}

type Game struct {
	noiseImage *image.RGBA
}

type Rect struct{
	x, y int
	width, heigth int
} 

func newRect(x int, y int, height int, width int) Rect{
	return Rect{
		x : x,
		y: y,
		heigth: height,
		width: width,
	}
}

func (g *Game)DrawBackground(screen *ebiten.Image){

	
	for x:= 0; x < screenWidth; x++{
		for y := 0; y < screenWidth ; y++{
			col := getColour(float64(y)/ screenHeight)
			screen.Set(x, y, col)

		}
	}
}
func (g *Game)DrawRect(screen *ebiten.Image){

	for x:= rect.x; x < rect.x + screenWidth; x++{
		for y := rect.y ; y < rect.y + screenHeight; y++{
			screen.Set(x, y, color.Black)

		}
	}
}

func (g *Game) Update() error {
	return nil
}

func getColour(i float64) color.Color{
	if i < 0.33{
		return red
	}else if i < 0.66{
		return blue
	} else {
		return green
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)

	purpleCol := color.RGBA{234,128,255, 255}

	for x := 100; x < 200; x++ {
		for y := 100; y < 200; y++ {
			screen.Set(x, y, purpleCol)
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebiten Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}