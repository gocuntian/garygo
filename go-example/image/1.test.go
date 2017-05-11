package main

import (
    "image"
    "image/color"
    "image/draw"
)

func main(){
    Color()
}

func Color(){
    c := color.RGBA{255,0,255,255}
    r := image.Rect(0,0,640,480)
    dst := image.NewRGBA(r)
    draw.Draw(dst, r, &image.Uniform{c}, image.ZP, draw.Src)
    m := image.NewRGBA(image.Rect(0,0,640,480))

    blue := color.RGBA{0,0,255,255}
    draw.Draw(m, m.Bounds(), &image.Uniform{blue},image.ZP, draw.Src)

    draw.Draw(m, m.Bounds(), image.Transparent, image.ZP,draw.Src)
}