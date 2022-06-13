package main

import (
	"image/color"
	"math"
)

type camera struct {
	worldUnitPosX, worldUnitPosY int
	worldPosX, worldPosY float64
	unitSize, mapSize    int
	angle                float64
	fov                  float64

	traversable bool
	transparent bool
	colour      color.Color
}
func NewCamera(x float64, y float64) camera {
	return camera{
		worldUnitPosX:   int(math.Floor(x)),
		worldUnitPosY:   int(math.Floor(y)),
		worldPosX:   x,
		worldPosY:   y,
		angle:       0,
		fov:         90,
		unitSize:    64,
		mapSize:     15,
		traversable: true,
		transparent: true,
		colour:      Green(),
	}
}

func (camera camera) GetColour() color.Color {
	return color.White
}

func (camera camera) GetCoord() (float64, float64) {
	return float64(camera.worldPosX), float64(camera.worldPosY)

}

func (camera camera) IsTraversable() bool {
	return camera.transparent
}

func (camera camera) IsTransparent() bool {
	return camera.transparent
}

func (camera camera) LineOfSightIntersect(float64, float64) bool {
	//Todo this is the calc to see if a line passes more into a square than not
	panic("cam LineOfSightIntersect Not implemented")
}

func (camera camera) GetCameraVector() (float64, float64, float64) {
	return camera.worldPosX, camera.worldPosY, camera.angle
}

func (camera *camera) UpdatePosition(newX float64, newY float64, newAngle float64) {
	camera.worldPosX = newX
	camera.worldPosY = newY
	camera.angle = newAngle
}
func (camera camera) GetSize() int{
	return camera.unitSize
}