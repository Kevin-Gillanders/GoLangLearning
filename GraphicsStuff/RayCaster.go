package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type rayCaster struct {
	rayCollection []ray
}

func NewRayCaster(fov int, rayWidth float64) rayCaster {
	rayCollection := make([]ray, fov/int(rayWidth)+1)
	return rayCaster{
		rayCollection: rayCollection,
	}

}

func (rayCaster *rayCaster) DrawRays(dst *ebiten.Image) {

	for _, rayToCast := range rayCaster.rayCollection {
		for _, op := range rayToCast.TransformRay() {
			if op == nil {
				continue
			}
			// Filter must be 'nearest' filter (default).
			// Linear filtering would make edges blurred.
			dst.DrawImage(emptySubImage, op)
		}
	}
}

func (rayCaster *rayCaster) UpdateRays(x, y, angle float64) {
	ceiling := color.Black
	wall := DeepRed()
	floor := color.White

	dX := -x 
	dY := -y 
	log.Println("x, y", x, y)
	log.Println("dx, dy", dX, dY)
	for idx, ray := range rayCaster.rayCollection {
		position := float64(idx) * rayWidth








		offset := (NormaliseFloat(float64(idx),
			float64(len(rayCaster.rayCollection))+1))
		startOfWall, endOfWall := ray.CalculateHeight(dist - dist*offset)

		ray = RayOutline(
			/*x1     :*/ position,
			/*y1     :*/ 0,
			/*x2     :*/ position,
			/*y2     :*/ float64(worldY),
			/*wall x :*/ startOfWall,
			/*wall y :*/ endOfWall,
			/*width  :*/ rayWidth,
			/*colour :*/ PickRandomColour(),
		)
		ray.SetColours(ceiling, wall, floor)
		rayCaster.rayCollection[idx] = ray
		// floorLevel += step
	}
}
