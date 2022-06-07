package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"strings"
	"time"

	// "math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var emptyImage *ebiten.Image = ebiten.NewImage(3, 3)
var emptySubImage *ebiten.Image = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var rayWidth, moveSpeed, rotationSpeed float64
var worldX, worldY int

var dist = 0.0
var inc bool = false
var step = 0.1

var t = time.Now().Add(time.Second * 1)

// Game implements ebiten.Game interface.
type Game struct {
    keys []ebiten.Key
}


func init(){
	worldX = 640
	worldY = 640
	emptyImage.Fill(color.White)

}


// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return worldX, worldY

}


func (g *Game) Update() error {
 
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	
	screen.Fill(color.RGBA{255, 255, 255, 255})


	randomOffset := rand.Float32() * 200


	var path vector.Path
	// H
	path.MoveTo(randomOffset + 20, 20 + randomOffset)
	path.LineTo(randomOffset + 20, 120 + randomOffset)
	path.LineTo(randomOffset + 30, 120 + randomOffset)
	path.LineTo(randomOffset + 30, 75 + randomOffset)
	path.LineTo(randomOffset + 110, 75 + randomOffset)
	path.LineTo(randomOffset + 110, 120 + randomOffset)
	path.LineTo(randomOffset + 120, 120 + randomOffset)
	path.LineTo(randomOffset + 120, 20 + randomOffset)
	path.LineTo(randomOffset + 110, 20 + randomOffset)
	path.LineTo(randomOffset + 110, 65 + randomOffset)
	path.LineTo(randomOffset + 30, 65 + randomOffset)
	path.LineTo(randomOffset + 30, 20 + randomOffset)
	path.LineTo(randomOffset + 20, 20 + randomOffset)
	
	//I
	path.MoveTo(randomOffset + 140, 20 + randomOffset)
	path.LineTo(randomOffset + 140, 20 + randomOffset)
	path.LineTo(randomOffset + 240, 20 + randomOffset)
	path.LineTo(randomOffset + 240, 30 + randomOffset)
	path.LineTo(randomOffset + 200, 30 + randomOffset)
	path.LineTo(randomOffset + 200, 110 + randomOffset)
	path.LineTo(randomOffset + 240, 110 + randomOffset)
	path.LineTo(randomOffset + 240, 120 + randomOffset)
	path.LineTo(randomOffset + 140, 120 + randomOffset)
	path.LineTo(randomOffset + 140, 110 + randomOffset)
	path.LineTo(randomOffset + 190, 110 + randomOffset)
	path.LineTo(randomOffset + 190, 30 + randomOffset)
	path.LineTo(randomOffset + 140, 30 + randomOffset)
	path.LineTo(randomOffset + 140, 20 + randomOffset)

	// Z
	// -
	path.MoveTo(randomOffset + 20, 140 + randomOffset)
	path.LineTo(randomOffset + 20, 140 + randomOffset)
	path.LineTo(randomOffset + 120, 140 + randomOffset)
	// |
	path.LineTo(randomOffset + 120, 150 + randomOffset)
	// /
	path.LineTo(randomOffset + 30, 230 + randomOffset)
	path.LineTo(randomOffset + 120, 230 + randomOffset)
	path.LineTo(randomOffset + 120, 240 + randomOffset)
	path.LineTo(randomOffset + 20, 240 + randomOffset)
	path.LineTo(randomOffset + 20, 230 + randomOffset)
	// /
	path.LineTo(randomOffset + 110, 150 + randomOffset)
	path.LineTo(randomOffset + 20, 150 + randomOffset)
	path.LineTo(randomOffset + 20, 140 + randomOffset)

	// O
	// -
	path.MoveTo(randomOffset + 140, 140 + randomOffset)
	path.LineTo(randomOffset + 140, 140 + randomOffset)
	path.LineTo(randomOffset + 240, 140 + randomOffset)
	path.LineTo(randomOffset + 240, 150 + randomOffset)
	path.LineTo(randomOffset + 140, 150 + randomOffset)
	path.LineTo(randomOffset + 140, 140 + randomOffset)
	path.LineTo(randomOffset + 140, 240 + randomOffset)
	path.LineTo(randomOffset + 150, 240 + randomOffset)
	path.LineTo(randomOffset + 150, 140 + randomOffset)
	path.LineTo(randomOffset + 140, 140 + randomOffset)
	path.MoveTo(randomOffset + 140, 240 + randomOffset)
	path.LineTo(randomOffset + 140, 240 + randomOffset)
	path.LineTo(randomOffset + 240, 240 + randomOffset)
	path.LineTo(randomOffset + 240, 230 + randomOffset)
	path.LineTo(randomOffset + 140, 230 + randomOffset)
	path.LineTo(randomOffset + 140, 240 + randomOffset)
	path.MoveTo(randomOffset + 240, 240 + randomOffset)
	path.LineTo(randomOffset + 240, 240 + randomOffset)
	path.LineTo(randomOffset + 230, 240 + randomOffset)
	path.LineTo(randomOffset + 230, 140 + randomOffset)
	path.LineTo(randomOffset + 240, 140 + randomOffset)
	path.LineTo(randomOffset + 240, 240 + randomOffset)


	path.MoveTo(randomOffset + 260, 140 + randomOffset)
	path.LineTo(randomOffset + 260, 140 + randomOffset)
	path.LineTo(randomOffset + 260, 240 + randomOffset)
	path.LineTo(randomOffset + 360, 240 + randomOffset)
	path.LineTo(randomOffset + 360, 230 + randomOffset)
	path.LineTo(randomOffset + 270, 230 + randomOffset)
	path.LineTo(randomOffset + 270, 195 + randomOffset)
	path.LineTo(randomOffset + 360, 195 + randomOffset)
	path.LineTo(randomOffset + 360, 185 + randomOffset)
	path.LineTo(randomOffset + 270, 185 + randomOffset)
	path.LineTo(randomOffset + 270, 185 + randomOffset)
	path.LineTo(randomOffset + 270, 150 + randomOffset)
	path.LineTo(randomOffset + 360, 150 + randomOffset)
	path.LineTo(randomOffset + 360, 140 + randomOffset)
	path.LineTo(randomOffset + 260, 140 + randomOffset)
	

	// path.LineTo(120, 140)
	// path.ArcTo(170, 250, 220, 190, 50)
	// path.QuadTo(220, 140, 120, 140)
	// path.QuadTo(150, 57.5, 100, 45)
	// path.QuadTo(150, 32.5, 100, 20)
	// path.LineTo(30, 60)
	// path.LineTo(30, 50)
	// path.LineTo(70, 50)
	// path.LineTo(70, 40)
	// path.LineTo(30, 40)
	// path.LineTo(30, 30)
	// path.LineTo(70, 30)
	// path.LineTo(70, 20)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}


	vertices, indexes := path.AppendVerticesAndIndicesForFilling(nil, nil)

	for i := range vertices{
		vertices[i].SrcX   = 1
		vertices[i].SrcY   = 1
		vertices[i].ColorR = 0xdb / float32(0xff)
		vertices[i].ColorG = 0x56 / float32(0xff)
		vertices[i].ColorB = 0x20 / float32(0xff)
	} 
	fmt.Println("vertex", vertices[0])
	fmt.Println("index"  , indexes[0])
	fmt.Println("================")

	screen.DrawTriangles(vertices, indexes, emptyImage, op)
	// c := color.RGBA{0, 0, 255, 255}
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(100, 100)
	// op.GeoM.Translate(10, 10 )
	// op.ColorM.ScaleWithColor(c)

	// screen.DrawImage(emptyImage, op)

	fps := fmt.Sprintf("FPS : %v", ebiten.CurrentFPS())

	if len(g.keys) == 0 {
		ebitenutil.DebugPrint(screen, fps)

	} else {
		keyStrs := []string{}
		for _, p := range g.keys {
			keyStrs = append(keyStrs, p.String())
		}
		ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))

	}
}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(worldY, worldX)
	ebiten.SetWindowTitle("Go Raycasting engine")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
