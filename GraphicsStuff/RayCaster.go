package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ray struct{
    x1, y1 float64
    x2, y2 float64

    width float64

    colour color.Color
}


func CreateARay(x1 float64, y1 float64, x2 float64, y2 float64, width float64, colour color.Color) *Ray{
    return &Ray{
        x1: x1,
        y1: y1,
        x2: x2,
        y2: y2,
        width: width,
        colour: colour,
    }
}


var emptyImage *ebiten.Image = ebiten.NewImage(3, 3)
var emptySubImage *ebiten.Image = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

var rotationDegrees, rotationSpeed, rayWidth float64
var worldX, worldY int

func init() {
    emptyImage.Fill(color.White)
    
    worldX = 640 
    worldY = 480
    
    rayWidth = 10

    rotationDegrees = 0.0 
    rotationSpeed = 0.1 


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

    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    // Write your game's rendering.

    for i := 0; i < worldX; i += int(rayWidth){

        var red, blue, green  uint8

        red = uint8(rand.Intn(255))
        blue = uint8(rand.Intn(255))
        green = uint8(rand.Intn(255))


        c := color.RGBA{red, blue, green, 255}

        ray := CreateARay(float64(i), 0, float64(i), float64(worldY) - 20, rayWidth, c)
        CastRay(screen, ray)
    }


    // ray1 := CreateARay(0, 0, 0, float64(worldY) - 20, 10, Blue())
    // ray2 := CreateARay(10, 0, 10, float64(worldY) - 40, 10, Green())
    // CastRay(screen, ray1)
    // CastRay(screen, ray2)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    
    return worldX, worldY

}

func CastRay(dst *ebiten.Image, rayToCast *Ray) {
    
    length := math.Sqrt(math.Pow((rayToCast.x2 - rayToCast.x1), 2) + math.Pow((rayToCast.y2 - rayToCast.y1), 2))
    fmt.Println("=============")
    fmt.Printf("X1 : %05f, Y1 : %05f\n", rayToCast.x1, rayToCast.y1)
    fmt.Printf("X2 : %05f, Y2 : %05f\n", rayToCast.x2, rayToCast.y2)
    fmt.Printf("Length %05f\n", length)
    fmt.Println("=============")


    op := &ebiten.DrawImageOptions{}
    
    op.GeoM.Scale(rayToCast.width, length )


    // op.GeoM.Rotate(DegreesToRadians(90))
    op.GeoM.Translate(rayToCast.x1, 0)
    // op.GeoM.Translate(0, y2)
    op.ColorM.ScaleWithColor(rayToCast.colour)

    // Filter must be 'nearest' filter (default).
    // Linear filtering would make edges blurred.
    dst.DrawImage(emptySubImage, op)
}



func main() {
    game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(worldX, worldY)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}