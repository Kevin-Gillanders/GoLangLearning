package main

import "image/color"

type emptyCell struct {
	worldX, worldY float64
	traversable    bool
	transparent    bool
	colour         color.Color
}

func NewEmptyCell(x float64, y float64) emptyCell {
	return emptyCell{
		worldX:      x,
		worldY:      y,
		traversable: true,
		transparent: true,
		colour:      White(),
	}
}

func (emptyCell emptyCell) GetColour() color.Color {
	return emptyCell.colour
}

func (emptyCell emptyCell) GetCoord() (float64, float64) {
	return float64(emptyCell.worldX), float64(emptyCell.worldY)

}

func (emptyCell emptyCell) IsTraversable() bool {
	return emptyCell.transparent
}

func (emptyCell emptyCell) IsTransparent() bool {
	return emptyCell.transparent
}
func (emptyCell emptyCell) LineOfSightIntersect(float64, float64) bool {
	//Todo this is the calc to see if a line passes more into a square than not
	panic("emp LineOfSightIntersect Not implemented")
}
