package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type ray struct {
	ceiling segment
	wall    segment
	floor   segment

	width float64

	colour color.Color
}

func RayOutline(x1 float64, y1 float64, x2 float64, y2 float64, wallTopY float64, wallBotY float64, width float64, entityCol color.Color) ray {

	return ray{
		ceiling: segment{
			x1:     x1,
			y1:     0,
			x2:     x2,
			y2:     wallTopY,
			colour: Orange(),
		},
		wall: segment{
			x1:     x1,
			y1:     wallTopY,
			x2:     x2,
			y2:     wallBotY,
			colour: Purple(),
		},
		floor: segment{
			x1:     x1,
			y1:     wallBotY,
			x2:     x2,
			y2:     float64(worldY),
			colour: Cyan(),
		},
		width: width,
	}
}

func (ray *ray) SetColours(ceilingColour color.Color, wallColour color.Color, floorColour color.Color) {
	ray.ceiling.colour = ceilingColour
	ray.wall.colour = wallColour
	ray.floor.colour = floorColour

}

func (rayToCast *ray) TransformRay() []*ebiten.DrawImageOptions {

	// var ceiling, wall, floor *ebiten.DrawImageOptions
	ceiling := rayToCast.ceiling.CreateTranformation(rayToCast.width)
	wall := rayToCast.wall.CreateTranformation(rayToCast.width)
	floor := rayToCast.floor.CreateTranformation(rayToCast.width)

	return []*ebiten.DrawImageOptions{ceiling, wall, floor}
}

func (ray *ray) CalculateHeight(distance float64) (float64, float64) {

	//Start the wall as being non existant
	var startOfWall, endOfWall float64

	startOfWall = float64(worldY) / 2
	endOfWall = float64(worldY) / 2

	vanishingPoint := 10.0

	// Get the normalised distance as a % of the culling point
	// eg if 10 distance need to the the diff of it against culling
	// so 10 - 10 = 0 normalised is 0.0
	// dist 0
	// 10 - 0 = 10 normalised is 1.0
	distaceNormalised := NormaliseFloat((vanishingPoint - distance), vanishingPoint)

	if distaceNormalised <= 0 {
		return startOfWall, endOfWall
	}

	// - moving the wall up
	startOfWall -= (float64(worldY) / 2) * distaceNormalised

	// + moving the wall down
	endOfWall += (float64(worldY) / 2) * distaceNormalised

	return startOfWall, endOfWall

}
