package main

import (
	"image/color"
	"math/rand"
)

func Black() color.Color {

	c := color.Black
	return c
}

func White() color.Color {

	c := color.White
	return c
}

func Blue() color.Color {

	c := color.RGBA{0, 0, 255, 255}
	return c
}

func LightBlue() color.Color {

	c := color.RGBA{52, 143, 235, 255}
	return c
}

func DarkBlue() color.Color {

	c := color.RGBA{22, 14, 140, 255}
	return c
}

func Cyan() color.Color {

	c := color.RGBA{52, 229, 235, 255}
	return c
}

func WeirdBlue() color.Color {

	c := color.NRGBA{0, 0, 255, 255}
	return c
}

func Red() color.Color {

	c := color.RGBA{255, 0, 0, 255}
	return c
}

func LightRed() color.Color {

	c := color.RGBA{232, 181, 167, 255}
	return c
}

func DeepRed() color.Color {

	c := color.RGBA{158, 13, 23, 255}
	return c
}

func Green() color.Color {

	c := color.RGBA{75, 207, 19, 255}
	return c
}

func LightGreen() color.Color {

	c := color.RGBA{167, 232, 175, 255}
	return c
}

func Yellow() color.Color {

	c := color.RGBA{207, 203, 19, 255}
	return c
}

func Purple() color.Color {

	c := color.RGBA{141, 13, 158, 255}
	return c
}

func LightPurple() color.Color {

	c := color.RGBA{142, 131, 168, 255}
	return c
}

func Orange() color.Color {

	c := color.RGBA{222, 131, 27, 255}
	return c
}

func PickRandomColour() color.Color {
	idx := rand.Intn(14)

	switch idx {
	case 0:
		return Orange()
	case 1:
		return LightPurple()
	case 2:
		return Purple()
	case 3:
		return Yellow()
	case 4:
		return LightGreen()
	case 5:
		return Green()
	case 6:
		return Red()
	case 7:
		return DeepRed()
	case 8:
		return Blue()
	case 9:
		return LightBlue()
	case 10:
		return DarkBlue()
	case 11:
		return Cyan()
	case 12:
		return WeirdBlue()
	case 13:
		return LightRed()
	default:
		return Purple()

	}
}
