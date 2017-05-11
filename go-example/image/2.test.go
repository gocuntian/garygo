package main

import (
    "fmt"
    "image"
)
//生成图像
func main(){
    myImage := image.NewRGBA(image.Rect(0,0,10,4))
    myImage.Pix[0] = 255
    myImage.Pix[1] = 0
    myImage.Pix[2] = 0
    myImage.Pix[3] = 255

    fmt.Println(myImage.Pix)
    fmt.Println(myImage.Stride)
}


// type RGBA struct {
//     // Pix保管图像的像素色彩信息，顺序为R, G, B, A
//     // 像素(x, y)起始位置是Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4]
//     Pix []uint8
//     // Stride是Pix中每行像素占用的字节数
//     Stride int
//     // Rect是图像的范围
//     Rect Rectangle
// }
// RGBA类型代表一幅内存中的图像，其At方法返回color.RGBA类型的值。

// func NewRGBA(r Rectangle) *RGBA
// NewRGBA函数创建并返回一个具有指定范围的RGBA。

// func Rect(x0, y0, x1, y1 int) Rectangle
// 返回一个矩形Rectangle{Pt(x0, y0), Pt(x1, y1)}。