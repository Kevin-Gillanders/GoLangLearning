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

func DistanceBetweenTwoPoints(x1 float64, y1 float64, x2 float64, y2 float64) float64{

	//Pythagoras theorm
	return math.Sqrt(math.Pow( x2 - x1 , 2) + math.Pow( y2 - y1 , 2) )
}