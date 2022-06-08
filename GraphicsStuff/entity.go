package main

import "image/color"

type entity interface {
	GetColour() color.Color
	GetCoord() (float64, float64)
	GetSize() int
	IsTransparent() bool
	IsTraversable() bool
	LineOfSightIntersect(float64, float64) bool
}
