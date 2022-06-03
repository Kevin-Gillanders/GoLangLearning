package main

import "image/color"


type wall struct{
	worldX, worldY float64
	traversable bool
	transparent bool
	colour color.Color

}

func NewWall(x float64, y float64) wall{
	return wall{
		worldX: x,
		worldY: y,
		traversable: false,
		transparent: false,
		colour: Red(),
	}
}

func (wall wall) GetColour() color.Color{
	return wall.colour
}

func (wall wall)GetCoord() (float64, float64){
	return float64(wall.worldX), float64(wall.worldY)

}

func (wall wall)IsTransparent() bool{
	return wall.transparent
}

func (wall wall)IsTraversable() bool{
	return wall.transparent
}

func (wall wall)LineOfSightIntersect(float64, float64) bool{
	//Todo this is the calc to see if a line passes more into a square than not
	panic("wall los Not implemented")
}
