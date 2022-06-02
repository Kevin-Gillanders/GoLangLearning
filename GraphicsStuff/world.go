package main

type world struct {
	x, y int
	camera camera
	entities [][] entity

	rotationDegrees, rotationSpeed float64

	rayCaster rayCaster

}

func CreateWorld(worldDefinition [][]rune, rayWidth float64, screenX int, screenY int) *world {
	entities := [][] entity{}
	var camera *camera
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
	return &world{
		entities: entities,
		x: len(entities[0]),
		y: len(entities),
		camera: *camera,
		rotationSpeed : 0.1,
		rayCaster: *rayCaster,
	}
}


func CastRay(originX float64, originY float64, theta float64) *ray {
	panic("world CastRay Not implemented")
	
}

func (world *world) RotateCamera(clockwise bool){
	if clockwise{
		world.camera.Rotate(world.rotationSpeed)
	}else{
		world.camera.Rotate(world.rotationSpeed * -1)
	}
}