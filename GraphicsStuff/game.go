package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	// "math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var emptyImage *ebiten.Image = ebiten.NewImage(3, 3)
var emptySubImage *ebiten.Image = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var rotationDegrees, rotationSpeed, rayWidth float64
var worldX, worldY int

var rayCollection []*ray
var updateOnce bool = false
var dist = 0.0
var inc bool = false
var step = 0.1

var World [][]rune

func init() {
	emptyImage.Fill(color.White)

	World = [][]rune{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '>', '-', '-', 'X'},
		{'X', '-', 'X', '-', 'X'},
		{'X', '-', '-', '-', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}

	worldX = 640
	worldY = 480

	rayWidth = 5

	rotationDegrees = 0.0
	rotationSpeed = 0.1

	// roof
	// wall
	// floor
	rayCollection = make([]*ray, worldX/int(rayWidth), worldX/int(rayWidth))
	UpdateRays()

}

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	// xCoor2 = xCoor2 + math.Cos(1)
	// yCoor2 = yCoor2 - math.Sin(1)

	// rotationDegrees = math.Mod(rotationDegrees + rotationSpeed, 360.0)
	if !updateOnce {
		UpdateRays()
	}

	dist = UpdateDist(dist, step)

	return nil
}

func UpdateDist(dist float64, step float64) float64 {
	if dist >= 10.0 {
		inc = false
	}
	if dist <= 0 {
		inc = true
	}
	if inc {
		return (dist + step)
	} else {
		return (dist - step)
	}
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	fps := fmt.Sprintf("FPS : %v", ebiten.CurrentFPS())

	for _, ray := range rayCollection {
		DrawRay(screen, ray)
	}
	ebitenutil.DebugPrint(screen, fps)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return worldX, worldY

}

func main() {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(worldX, worldY)
	ebiten.SetWindowTitle("Go Raycasting engine")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
