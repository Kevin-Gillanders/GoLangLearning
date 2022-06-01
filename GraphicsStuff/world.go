package main

type world struct {
	x, y int
	camera camera
	entities [][] entity
}

func CreateWorld(worldDefinition [][]rune) *world {
	entities := [][] entity{}
	var camera camera
	for y, row := range worldDefinition{
		entityRow := [] entity{}
		for x, letter := range row{

			switch(letter){
				case('X'):
					entityRow = append(entityRow, NewWall(float64(x), float64(y)))
				
			}
		} 
		entities = append(entities, entityRow)
	}
	return &world{
		entities: entities,
		x: len(entities[0]),
		y: len(entities),
		camera: camera,
	}
}


func CastRay(originX float64, originY float64, theta float64) *ray {
	panic("Not implemented")
	
}