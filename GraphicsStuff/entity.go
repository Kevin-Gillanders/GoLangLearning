package main

import "image/color"


type entity interface {
	GetColour() color.Color
	GetCoord() (float64, float64)
	IsTransparent() bool
	LineOfSightIntersect(float64, float64) bool
}
