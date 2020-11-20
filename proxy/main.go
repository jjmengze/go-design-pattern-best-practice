package main

import "fmt"

func main() {
	var image, imageAudit imageShower
	image = NewRealImage("realImage")        // loadFileFromDisk() immediately
	imageAudit = NewProxyImage("proxyImage") // loadFileFromDisk() deferred
	// file is already loaded and display gets called directly
	image.ShowImage()
	// load file from disk
	// and then forward display call to the real image
	fmt.Println("I want to use image , but user have to be audit so let talk to imageAudit proxy !")
	imageAudit.ShowImage()
	fmt.Println("thanks you imageAudit proxy !")
}
