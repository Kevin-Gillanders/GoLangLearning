package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Segment struct {
	x1, y1 float64
	x2, y2 float64
	colour color.Color
}

func (seg *Segment) CreateTranformation(width float64) *ebiten.DrawImageOptions {

	length := math.Sqrt(math.Pow((seg.x2-seg.x1), 2) + math.Pow((seg.y2-seg.y1), 2))
	// fmt.Println("=============")
	// fmt.Printf("X1 : %05f, Y1 : %05f\n", seg.x1, seg.y1)
	// fmt.Printf("X2 : %05f, Y2 : %05f\n", seg.x2, seg.y2)
	// fmt.Printf("Length %05f, width %05f\n", length, width)
	// fmt.Println("=============")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(width, length)

	// op.GeoM.Rotate(DegreesToRadians(90))
	op.GeoM.Translate(seg.x1, seg.y1)
	// op.GeoM.Translate(0, y2)
	op.ColorM.ScaleWithColor(seg.colour)
	return op
}
