package main

import (
	"fmt"
	"image/color"
	"reflect"

	// "image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var square *ebiten.Image
var cameraImage *ebiten.Image
var line *ebiten.Image

var lineX, lineY int

type world struct {
	x, y     int
	camera   camera
	entities [][]entity

	movementSpeed, rotationSpeed float64

	rayCaster rayCaster
}

func CreateWorld(worldDefinition [][]rune, rayWidth float64, screenX int, screenY int, moveSpeed float64, rotationSpeed float64) world {
	fmt.Println("Creating world")

	lineX = 20
	lineY = 1

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
	rayCaster := NewRayCaster(int(camera.fov), rayWidth)
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
	
	newX := currentX
	newY := currentY
	newTheta = currentTheta

	for _, k := range keys {
		// <- : Angle of rotation - rotation speed
		// -> : Angle of rotation + rotation speed

		// ^  : -Y
		// v  : +Y

		// A : x - MS
		// D : x + MS

		// W : Y - MS
		// S : Y + MS
		switch k {
			//Grid controls
			case ebiten.KeyA:
				newX = currentX - moveSpeed
				newY = currentY
			
			case ebiten.KeyW:
				newX = currentX 
				newY = currentY - moveSpeed
			
			case ebiten.KeyS:
				newX = currentX 
				newY = currentY + moveSpeed
			
			case ebiten.KeyD:
				newX = currentX + moveSpeed
				newY = currentY 
			
			case ebiten.KeyQ:
				newTheta = (currentTheta - world.rotationSpeed)
				fmt.Printf("Left <= currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
				if newTheta < 0 {
					newTheta = newTheta + 360
				}
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)

			case ebiten.KeyE:
				newTheta = math.Mod(currentTheta+world.rotationSpeed, 360)
				fmt.Printf("Right => currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)
			



			//Tank controls
			case ebiten.KeyArrowLeft:
				newTheta = (currentTheta - world.rotationSpeed)
				fmt.Printf("Left <= currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
				if newTheta < 0 {
					newTheta = newTheta + 360
				}
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)

			case ebiten.KeyArrowUp:
				distance = distance + world.movementSpeed
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)
				
			case ebiten.KeyArrowDown:
				distance = distance - world.movementSpeed
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)
			
			case ebiten.KeyArrowRight:
				newTheta = math.Mod(currentTheta+world.rotationSpeed, 360)
				fmt.Printf("Right => currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
				newX, newY = DerivedNewPoint(currentX, currentY, distance, newTheta)
			
			case ebiten.KeyR:
				world.camera.UpdatePosition(1, 1, 0)
				return world
			
			case ebiten.KeySpace:
				newTheta = math.Mod(currentTheta+90, 360)
			
			case ebiten.KeyL:
				newTheta = math.Mod(currentTheta+45, 360)
			
			case ebiten.KeyZ:
				lineX ++
				line = ebiten.NewImage(lineX, lineY)
				line.Fill(color.White)
			
			case ebiten.KeyX:
				if lineX > 1 {
					lineX --
					line = ebiten.NewImage(lineX, lineY)
					line.Fill(color.White)
				}
		}
	}

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

	screen.Fill(Black()) //NRGBA{0xff, 0x00, 0x00, 0xff})

	squareX := worldX / len(world.entities)
	squareY := worldY / len(world.entities)

	if square == nil {
		square = ebiten.NewImage(squareX - 1, squareY - 1)
		square.Fill(color.White)
	}

	for iY, y := range world.entities {
		for iX, x := range y {
			// Fill the screen with #FF0000 color
			// Fill the square with the white color
			// fmt.Println(reflect.TypeOf(x))
			// fmt.Println(float64(iX) * float64(x.GetSize()), float64(iY) * float64(x.GetSize()))
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(iX) * float64(x.GetSize()), float64(iY) * float64(x.GetSize()))	
			op.ColorM.ScaleWithColor(x.GetColour())
			// Create an empty option struct
			// fmt.Println(len(world.entities), len(world.entities))

			// Draw the square image to the screen with an empty option
			screen.DrawImage(square, op)
		}
	}

	world.DrawCamera(screen)
	


}


func (world world) DrawCamera(screen *ebiten.Image){

	if cameraImage == nil || line == nil{
		line = ebiten.NewImage(lineX, lineY)
		line.Fill(color.White)
		cameraImage = ebiten.NewImage(world.camera.mapSize, world.camera.mapSize)
		cameraImage.Fill(color.White)
	}

	camX, camY := world.camera.GetCoord()
	op := &ebiten.DrawImageOptions{}



	log.Println(reflect.TypeOf(world.camera))
	log.Println(camX, camY)

	op.GeoM.Translate(camX, camY)
	op.ColorM.ScaleWithColor(world.camera.colour)
	screen.DrawImage(cameraImage, op)


	for i, ray := range world.rayCaster.rayCollection {

		op = &ebiten.DrawImageOptions{}

		op.GeoM.Rotate(DegreesToRadians(world.camera.angle + (float64(i) * ray.width)))

		op.GeoM.Translate(camX+(float64(world.camera.mapSize)/2), camY+(float64(world.camera.mapSize)/2))
		op.ColorM.ScaleWithColor(Black())

		screen.DrawImage(line, op)
	}

}
