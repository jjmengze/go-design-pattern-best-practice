package main

import (
	"fmt"
)

type image struct {
	fileName string
}

func (i *image) ShowImage() {
	fmt.Println("Displaying ", i.fileName)
}

func NewRealImage(filename string) *image {
	image := &image{fileName: filename}
	return image
}
func (i *image) LoadImageFromURL(url string) {
	fmt.Println("Load from URL: ", i.fileName)
}
