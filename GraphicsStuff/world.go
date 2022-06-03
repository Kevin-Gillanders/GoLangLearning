package main

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type world struct {
	x, y int
	camera camera
	entities [][] entity

	movementSpeed, rotationSpeed float64

	rayCaster rayCaster

}

func CreateWorld(worldDefinition [][]rune, rayWidth float64, screenX int, screenY int, moveSpeed float64, rotationSpeed float64) world {
	fmt.Println("Creating world")
	entities := [][] entity{}
	var camera camera
	for y, row := range worldDefinition{
		entityRow := [] entity{}
		for x, letter := range row{

			switch(letter){
				case('X'):
					entityRow = append(entityRow, NewWall(float64(x), float64(y)))
				case('-'):
					entityRow = append(entityRow, NewEmptyCell(float64(x), float64(y)))
				case('>'):
					camera = NewCamera(float64(x), float64(y))
					entityRow = append(entityRow, camera)
					
			}
		} 
		entities = append(entities, entityRow)
	}
	rayCaster := NewRayCaster(screenX, screenY, rayWidth)
	return world{
		entities: entities,
		x: len(entities[0]),
		y: len(entities),
		camera: camera,
		rotationSpeed : rotationSpeed,
		movementSpeed : moveSpeed,
		rayCaster: rayCaster,
	}
}


func CastRay(originX float64, originY float64, theta float64) ray {
	panic("world CastRay Not implemented")
	
}

func (world world) UpdateCameraPosition(keys []ebiten.Key) world{

	//https://math.stackexchange.com/questions/143932/calculate-point-given-x-y-angle-and-distance

	currentX, currentY, currentTheta := world.camera.GetCameraVector()
	var distance, newTheta float64 

	newTheta = currentTheta

	for _, k := range keys{
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
				newTheta = math.Mod(currentTheta + world.rotationSpeed, 360)
				fmt.Printf("Right => currentTheta : %v NewTheta %v\n", currentTheta, newTheta)
			case ebiten.KeyR:
				world.camera.UpdatePosition(1, 1, 0)
				return world
			case ebiten.KeySpace:
				newTheta = math.Mod(currentTheta + 90, 360)
			case ebiten.KeyL:
				newTheta = math.Mod(currentTheta + 45, 360)

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