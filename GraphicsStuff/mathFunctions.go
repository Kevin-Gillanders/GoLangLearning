package main

import "math"

func DegreesToRadians(degrees float64) float64 {

	radians := degrees * (math.Pi / 180)

	return radians

}

func RadiansToDegrees(radians float64) float64 {

	degrees := radians * (180 / math.Pi)

	return degrees

}

func NormaliseFloat(numberToNormalise float64, normalisedTo float64) float64 {

	return float64(numberToNormalise) / float64(normalisedTo)
}
