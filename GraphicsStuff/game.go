package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"
	"time"

	// "math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var emptyImage *ebiten.Image = ebiten.NewImage(3, 3)
var emptySubImage *ebiten.Image = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var rayWidth, moveSpeed, rotationSpeed float64
var worldX, worldY int

var rayCollection []ray
var dist = 0.0
var inc bool = false
var step = 0.1

var World world
var t = time.Now().Add(time.Second * 1)

func init() {
	emptyImage.Fill(color.White)

	worldX = 640
	worldY = 480

	rayWidth = 20
	moveSpeed = 1
	rotationSpeed = 1

	worldOutline := [][]rune{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '>', '-', '-', 'X'},
		{'X', '-', 'X', '-', 'X'},
		{'X', '-', '-', '-', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}

	World = CreateWorld(
		worldOutline,
		rayWidth,
		worldX,
		worldY,
		moveSpeed,
		rotationSpeed)

	rayCollection = make([]ray, worldX/int(rayWidth), worldX/int(rayWidth))
	World.rayCaster.UpdateRays()

}

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if t.Before(time.Now()) {
		World = World.UpdateCameraPosition(g.keys)
		t = time.Now().Add(time.Second)
	}
	World.rayCaster.UpdateRays()

	dist = UpdateDist(dist, step)

	return nil
}

func UpdateDist(dist float64, step float64) float64 {
	//TODO Remove this debug
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

	World.rayCaster.DrawRays(screen)

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
