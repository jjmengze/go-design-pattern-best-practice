package main

import "fmt"

type VerticalScrollBars struct {
	*WindowDecorator
}

func NewVerticalScrollBars(window Window) *VerticalScrollBars {
	vsb := new(VerticalScrollBars)
	vsb.WindowDecorator = NewWindowDecorator(window)
	return vsb
}
func (vsb *VerticalScrollBars) Draw() {
	vsb.drawVerticalScrollBar()
	vsb.window.Draw()
}
func (vsb *VerticalScrollBars) drawVerticalScrollBar() {
	fmt.Println("draw a vertical scrollbar")
}
func (vsb *VerticalScrollBars) GetDescription() string {
	return vsb.window.GetDescription() + ", including vertical scrollbars"
}
