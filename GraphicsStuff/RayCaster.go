package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func DrawRay(dst *ebiten.Image, rayToCast *Ray) {

	for _, op := range rayToCast.TransformRay() {
		if op == nil {
			continue
		}
		// Filter must be 'nearest' filter (default).
		// Linear filtering would make edges blurred.
		dst.DrawImage(emptySubImage, op)
	}
}

func UpdateRays() {
	ceiling := color.Black
	wall := DeepRed()
	floor := color.White

	// wallWidth := 20.0
	// floorLevel := 0.0
	// step := (float64(worldY) - floorLevel - wallWidth) / (float64(worldX) / rayWidth)
	// step := 0.01
	for idx, ray := range rayCollection {
		position := float64(idx) * rayWidth

		// var red, blue, green  uint8

		// red = uint8(rand.Intn(255))
		// blue = uint8(rand.Intn(255))
		// green = uint8(rand.Intn(255))

		// c := color.RGBA{red, blue, green, 255}
		// upperBound := float64(outerIdx * (worldY / 3))

		// var top, bot float64

		//Normalise top and bottom to between 0 .. 1
		// top := rand.Intn(int(float64(worldY) * 0.5))

		// bot := rand.Intn(worldY - top) + top

		// bot := (floorLevel + step) + wallWidth
		// top := floorLevel - wallWidth

		// if float64(bot)/float64(worldY) > 1 {
		// 	panic("World bounds exceeded")
		// }
		// fmt.Printf("worldY : %v\n", worldY)
		// fmt.Printf("           min : %v ==== max : %v\n", top, bot)
		// fmt.Printf("Normalised min : %05f ==== max : %05f\n", float64(top)/float64(worldY), float64(bot)/float64(worldY))
		offset := (NormaliseFloat(float64(idx), float64(len(rayCollection)) + 1))
		startOfWall, endOfWall := ray.CalculateHeight(dist - dist * offset)


		ray = CastRay(
			/*x1     :*/ position,
			/*y1     :*/ 0,
			/*x2     :*/ position,
			/*y2     :*/ float64(worldY),
			/*wall x :*/ startOfWall,
			/*wall y :*/ endOfWall,
			/*width  :*/ rayWidth,
			/*colour :*/ false,
		)
		ray.SetColours(ceiling, wall, floor)
		rayCollection[idx] = ray
		// floorLevel += step
	}
}

