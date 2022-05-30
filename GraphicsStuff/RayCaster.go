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

var rayCollection [] *Ray



func init() {
    emptyImage.Fill(color.White)
    
    worldX = 640 
    worldY = 480
    
    rayWidth = 5

    rotationDegrees = 0.0 
    rotationSpeed = 0.1 

    // roof
    // wall
    // floor
    rayCollection = make([]*Ray, worldX / int(rayWidth), worldX / int(rayWidth))

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

    UpdateRays()

    return nil
}


// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    // Write your game's rendering.
    fps := fmt.Sprintf("FPS : %v", ebiten.CurrentFPS())
    
    for _, ray := range rayCollection{ 
        DrawRay(screen, ray)
    }
    ebitenutil.DebugPrint(screen, fps)   
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    
    return worldX, worldY

}

func DrawRay(dst *ebiten.Image, rayToCast *Ray) {
    
    for _, op := range rayToCast.TransformRay() {
        if op == nil{
            continue
        }
        // Filter must be 'nearest' filter (default).
        // Linear filtering would make edges blurred.
        dst.DrawImage(emptySubImage, op)
    }
}


func UpdateRays(){

    for idx, ray := range rayCollection{
        position := float64(idx) * rayWidth

        // var red, blue, green  uint8

        // red = uint8(rand.Intn(255))
        // blue = uint8(rand.Intn(255))
        // green = uint8(rand.Intn(255))


        // c := color.RGBA{red, blue, green, 255}
        // upperBound := float64(outerIdx * (worldY / 3))

        ray = CastRay(
                        /*x1     :*/ position, 
                        /*y1     :*/ 0, 
                        /*x2     :*/ position, 
                        /*y2     :*/ float64(worldY), 
                        /*wall x :*/ float64(worldY) * 0.33, 
                        /*wall y :*/ float64(worldY) * 0.66, 
                        /*width  :*/ rayWidth, 
                        /*colour :*/ PickRandomColour(),
                        )

        rayCollection[idx] = ray
    }

    
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