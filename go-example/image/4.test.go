package main

import (
    "fmt"
	"image"
	"image/png"
	"os"
    "log"
)
//从文件中读取图像
func main(){
    existingImageFile, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer existingImageFile.Close()

	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

    existingImageFile.Seek(0,0)

    loadedImage, err := png.Decode(existingImageFile)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(loadedImage)
}