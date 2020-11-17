package main


type SimpleWindow struct {}
func (window *SimpleWindow) Draw() {
	//draw a simple window
}
func (window *SimpleWindow) GetDescription() string {
	return "a simple window"
}