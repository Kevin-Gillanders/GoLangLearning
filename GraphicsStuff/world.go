package main

import (
	"fmt"
	"image/color"

	// "image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var square *ebiten.Image
var line *ebiten.Image

type world struct {
	x, y     int
	camera   camera
	entities [][]entity

	movementSpeed, rotationSpeed float64

	rayCaster rayCaster
}

func CreateWorld(worldDefinition [][]rune, rayWidth float64, screenX int, screenY int, moveSpeed float64, rotationSpeed float64) world {
	fmt.Println("Creating world")
	entities := [][]entity{}
	var camera camera
	for y, row := range worldDefinition {
		entityRow := []entity{}
		for x, letter := range row {

			switch letter {
			case ('X'):
				entityRow = append(entityRow, NewWall(float64(x), float64(y)))
			case ('-'):
				entityRow = append(entityRow, NewEmptyCell(float64(x), float64(y)))
			case ('>'):
				camera = NewCamera(float64(x), float64(y))
				entityRow = append(entityRow, camera)

			}
		}
		entities = append(entities, entityRow)
	}
	rayCaster := NewRayCaster(screenX, screenY, rayWidth)
	return world{
		entities:      entities,
		x:             len(entities[0]),
		y:             len(entities),
		camera:        camera,
		rotationSpeed: rotationSpeed,
		movementSpeed: moveSpeed,
		rayCaster:     rayCaster,
	}
}

func CastRay(originX float64, originY float64, theta float64) ray {
	panic("world CastRay Not implemented")

}

func (world world) UpdateCameraPosition(keys []ebiten.Key) world {

	//https://math.stackexchange.com/questions/143932/calculate-point-given-x-y-angle-and-distance

	currentX, currentY, currentTheta := world.camera.GetCameraVector()
	var distance, newTheta float64

	newTheta = currentTheta

	for _, k := range keys {
		//Todo add strafing if shift is held
		// A, <- : Angle of rotation - rotation speed
		// D, -> : Angle of rotation + rotation speed

		// W, ^  : -Y
		// S, v  : +Y
		//TODO Debug this as the behaviour is incorrect
		switch k {
		case ebiten.KeyA:
		case ebiten.KeyArrowLeft:
			newTheta = (currentTheta - world.rotationSpeed)
			fmt.Printf("Left <= currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
			if newTheta < 0 {
				newTheta = newTheta + 360
			}
		case ebiten.KeyW:
		case ebiten.KeyArrowUp:
			distance = distance + world.movementSpeed
		case ebiten.KeyS:
		case ebiten.KeyArrowDown:
			distance = distance - world.movementSpeed
		case ebiten.KeyD:
		case ebiten.KeyArrowRight:
			newTheta = math.Mod(currentTheta+world.rotationSpeed, 360)
			fmt.Printf("Right => currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
		case ebiten.KeyR:
			world.camera.UpdatePosition(1, 1, 0)
			return world
		case ebiten.KeySpace:
			newTheta = math.Mod(currentTheta+90, 360)
		case ebiten.KeyL:
			newTheta = math.Mod(currentTheta+45, 360)

		}
	}
	newX, newY := DerivedNewPoint(currentX, currentY, distance, newTheta)

	world.camera.UpdatePosition(newX, newY, newTheta)

	fmt.Println("================")
	fmt.Printf("currentX %v currentY %v angle %v distance %v\n", currentX, currentY, currentTheta, distance)
	fmt.Printf("newX %v newY %v newangle %v \n", newX, newY, newTheta)
	fmt.Println("================")

	// if clockwise{
	// 	world.camera.Rotate(world.rotationSpeed)
	// }else{
	// 	world.camera.Rotate(world.rotationSpeed * -1)
	// }
	return world
}

func (world world) Draw2DWorld(screen *ebiten.Image) {

	screen.Fill(Green()) //NRGBA{0xff, 0x00, 0x00, 0xff})

	squareX := worldX / len(world.entities)
	squareY := worldY / len(world.entities)

	if square == nil {
		square = ebiten.NewImage(squareX, squareY)
		square.Fill(color.White)
	}

	for iY, y := range world.entities {
		for iX, x := range y {
			// Fill the screen with #FF0000 color
			// Fill the square with the white color
			fmt.Println(x)
			// Create an empty option struct
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(iX*squareX), float64(iY*squareY))
			op.ColorM.ScaleWithColor(x.GetColour())
			// fmt.Println(len(world.entities), len(world.entities))

			// Draw the square image to the screen with an empty option
			screen.DrawImage(square, op)
		}
	}
}

func (world world) DrawGrid(screen *ebiten.Image) {

	squareY := worldY / len(world.entities)
	squareX := worldX / len(world.entities[0])

	if line == nil {
		line = ebiten.NewImage(2, worldX)
		line.Fill(color.White)
	}

	for x := 0; x <= worldX/squareX; x++ {

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(float64(x*squareX), 0)
		op.ColorM.ScaleWithColor(Black())

		// Draw the line image to the screen with an empty option
		screen.DrawImage(line, op)
	}

	line = ebiten.NewImage(2, worldY)
	line.Fill(color.White)

	for y := 0; y <= worldY/squareY; y++ {
		fmt.Println("y", y)
		fmt.Println("squareY", squareY)
		fmt.Println("worldY", worldY)
		fmt.Println("worldY / squareY", worldY/squareY)
		fmt.Println("================")
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(float64(y*squareY), 0)
		op.ColorM.ScaleWithColor(Black())

		// Draw the line image to the screen with an empty option
		screen.DrawImage(line, op)
	}

	// for iY, y := range world.entities {
	// 	for iX := range y {
	//         // Create an 16x16 image
	//         line = ebiten.NewImage(2, 2)
	// 	    //Horizontal
	//         line = ebiten.NewImage(2, squareX)
	// 	    line.Fill(color.Black)

	// 	    op := &ebiten.DrawImageOptions{}

	//    	    op.GeoM.Translate(float64(iX * squareX), float64(iY * squareY))

	// 	    fmt.Println(float64(iX * squareX), float64(iY * squareY))

	// 	    // Draw the square image to the screen with an empty option
	// 	    screen.DrawImage(line, op)

	// 	    //Vertical
	//         line = ebiten.NewImage(squareY, 2)
	// 	    line.Fill(color.Black)

	// 	    op = &ebiten.DrawImageOptions{}

	//    	    op.GeoM.Translate(float64(iX * squareX), float64(iY * squareY))

	// 	    fmt.Println(float64(iX * squareX), float64(iY * squareY))

	// 	    // Draw the square image to the screen with an empty option
	// 	    screen.DrawImage(line, op)

	// 	}
	// }
}
