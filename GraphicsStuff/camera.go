package main

import (
	"image/color"
	"math"
)

type camera struct {
	worldX, worldY float64
	angle float64
	fov float64

	traversable bool
	transparent bool
	colour color.Color
}


func NewCamera(x float64, y float64) *camera{
	return &camera{
		worldX: x,
		worldY: y,
		traversable: true,
		transparent: true,
		colour: White(),
	}
}

func (camera *camera) GetColour() color.Color{
	return camera.colour
}

func (camera *camera) GetCoord() (float64, float64){
	return float64(camera.worldX), float64(camera.worldY)

}

func (camera *camera) IsTraversable() bool{
	return camera.transparent
}

func (camera *camera) IsTransparent() bool{
	return camera.transparent
}

func (camera *camera) LineOfSightIntersect(float64, float64) bool{
	//Todo this is the calc to see if a line passes more into a square than not
	panic("cam LineOfSightIntersect Not implemented")
}

func(camera *camera) Rotate(angleOfRotation float64){
	camera.angle = math.Mod(camera.angle + angleOfRotation, 360)
}
