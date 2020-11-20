package main

import "fmt"

type imageShower interface {
	ShowImage()
}

type imageAudit struct {
	filename string
	*image
}

func NewProxyImage(filename string) *imageAudit {
	return &imageAudit{filename: filename}
}

func (ip *imageAudit) ShowImage() {
	if ip.image == nil {
		ip.image = NewRealImage(ip.filename)
	}
	// do something different origin (image struct)
	fmt.Print("I am proxy , if you want to use image you have to talk to me ...")
	ip.image.ShowImage()
	fmt.Print("image done , proxy will done other thing ....")
	// do so,ething
}
