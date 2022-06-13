package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"


	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var emptyImage *ebiten.Image = ebiten.NewImage(3, 3)
var emptySubImage *ebiten.Image = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var rayWidth, moveSpeed, rotationSpeed float64
var worldX, worldY int

var dist = 0.0
var inc bool = false
var step = 0.1

var World world
var t = time.Now().Add(time.Second * 1)

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

func init() {
	log.SetFlags(log.Lshortfile)
	emptyImage.Fill(color.White)

	worldX = 640
	worldY = 640

	rayWidth = 10
	moveSpeed = 5
	rotationSpeed = 5

	worldOutline := [][]rune{
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '>', '-', '-', '-', '-', '-', '-', '-', 'X'},
		{'X', '-', 'X', '-', '-', '-', '-', 'X', '-', 'X'},
		{'X', '-', '-', '-', '-', '-', '-', '-', '-', 'X'},
		{'X', 'X', 'X', 'X', '-', '-', 'X', 'X', 'X', '-'},
		{'X', 'X', 'X', 'X', '-', '-', 'X', 'X', 'X', 'X'},
		{'X', '-', '-', '-', '-', '-', '-', '-', '-', 'X'},
		{'X', '-', 'X', '-', '-', '-', '-', 'X', '-', 'X'},
		{'X', '-', '-', '-', '-', '-', '-', '-', '-', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', '-'},

	}

	World = CreateWorld(
		worldOutline,
		rayWidth,
		worldX,
		worldY,
		moveSpeed,
		rotationSpeed)

	// World.rayCaster.UpdateRays()

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	// if t.Before(time.Now()) {
	World = World.UpdateCameraPosition(g.keys)
	// 	t = time.Now().Add(time.Second)
	// }
	World.rayCaster.UpdateRays(World.camera.worldPosX, World.camera.worldPosY, World.camera.angle)

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
	World.Draw2DWorld(screen)

	// World.rayCaster.DrawRays(screen)
	// World.DrawGrid(screen)
	// if len(g.keys) == 0 {
	x, y := World.camera.GetCoord()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v\nWorldx : %v Worldy : %v", fps, x, y))

	// } else {
	// 	keyStrs := []string{}
	// 	for _, p := range g.keys {
	// 		keyStrs = append(keyStrs, p.String())
	// 	}
	// 	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))

	// }

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return worldX, worldY

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
