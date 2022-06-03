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

func DerivedNewPoint(x float64, y float64, distance float64, theta float64) (float64, float64){

	//https://math.stackexchange.com/questions/143932/calculate-point-given-x-y-angle-and-distance
	//https://math.stackexchange.com/questions/604324/find-a-point-n-distance-away-from-a-specified-point-in-a-given-direction

	theta = DegreesToRadians(theta)

	newX := x + (distance * math.Cos(theta))
	newY := y + (distance * math.Sin(theta))

	return newX, newY
}