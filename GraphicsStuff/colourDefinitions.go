package main

import "image/color"

func Blue() color.Color{

	c := color.RGBA{0, 0, 255, 255}
	return c
}

func WeirdBlue() color.Color{

	c := color.NRGBA{0, 0, 255, 255}
	return c
}

func Red() color.Color{

	c := color.RGBA{255, 0, 0, 255}
	return c
}

func Green() color.Color{

	c := color.RGBA{0, 255, 0, 255}
	return c
}