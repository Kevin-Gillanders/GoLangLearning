package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ray struct{
    x1, y1 float64
    x2, y2 float64

    ceiling Segment
    wall Segment
    floor Segment

    width float64

    colour color.Color
}
                     

func CastRay(x1 float64, y1 float64, x2 float64, y2 float64, wallTopY float64, wallBotY float64, width float64, colour color.Color) *Ray{
    return &Ray{
        x1: x1,
        y1: y1,
        x2: x2,
        y2: y2,
        ceiling: Segment {
            x1 : x1,
            y1 : 0,
            x2 : x2,
            y2 : wallTopY,
            colour: PickRandomColour(),
        },
        wall: Segment {
            x1 : x1,
            y1 : wallTopY,
            x2 : x2,
            y2 : wallBotY,
            colour: PickRandomColour(),
        },
        floor: Segment {
            x1 : x1,
            y1 : wallBotY,
            x2 : x2,
            y2 : 0,
            colour: PickRandomColour(),
        },
        width: width,
        colour: colour,
    }
}


func(rayToCast *Ray) TransformRay() []*ebiten.DrawImageOptions {
	

    // var ceiling, wall, floor *ebiten.DrawImageOptions
    ceiling := rayToCast.ceiling.CreateTranformation(rayToCast.width) 
    wall    := rayToCast.wall.CreateTranformation(rayToCast.width) 
    floor   := rayToCast.floor.CreateTranformation(rayToCast.width) 

    return [] *ebiten.DrawImageOptions{ ceiling, wall, floor}
}

