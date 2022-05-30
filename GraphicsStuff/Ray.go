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
                     

func CastRay(x1 float64, y1 float64, x2 float64, y2 float64, wallTopY float64, wallBotY float64, width float64, randomCol bool) *Ray{
    
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
            colour: Orange(),
        },
        wall: Segment {
            x1 : x1,
            y1 : wallTopY,
            x2 : x2,
            y2 : wallBotY,
            colour: Purple(),
        },
        floor: Segment {
            x1 : x1,
            y1 : wallBotY,
            x2 : x2,
            y2 : float64(worldY),
            colour: Cyan(),
        },
        width: width,
    }
}


func(ray *Ray) SetColours(ceilingColour color.Color, wallColour color.Color, floorColour color.Color) {
    ray.ceiling.colour = ceilingColour
    ray.wall.colour    = wallColour
    ray.floor.colour   = floorColour

}



func(rayToCast *Ray) TransformRay() []*ebiten.DrawImageOptions {
	
    // var ceiling, wall, floor *ebiten.DrawImageOptions
    ceiling := rayToCast.ceiling.CreateTranformation(rayToCast.width) 
    wall    := rayToCast.wall.CreateTranformation(rayToCast.width) 
    floor   := rayToCast.floor.CreateTranformation(rayToCast.width) 

    return [] *ebiten.DrawImageOptions{ ceiling, wall, floor}
}


func(ray *Ray) CalculateHeight( distance float64) (float64, float64) {
    
    //Start the wall as being non existant
    var startOfWall, endOfWall float64

    startOfWall = float64(worldY) / 2
    endOfWall = float64(worldY) / 2


    vanishingPoint := 10.0

    // Get the normalised distance as a % of the culling point
    // eg if 10 distance need to the the diff of it against culling
    // so 10 - 10 = 0 normalised is 0.0
    // dist 0
    // 10 - 0 = 10 normalised is 1.0
    distaceNormalised := NormaliseFloat((vanishingPoint - distance), vanishingPoint)

    if distaceNormalised <= 0 {
        return startOfWall, endOfWall
    }

    // - moving the wall up
    startOfWall -= (float64(worldY) / 2) * distaceNormalised

    // + moving the wall down
    endOfWall += (float64(worldY) / 2 ) * distaceNormalised

    return startOfWall, endOfWall


}