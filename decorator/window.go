package main

type Window interface {
	Draw()
	GetDescription() string
}

type WindowDecorator struct {
	window Window
}

func NewWindowDecorator(window Window) *WindowDecorator {
	return &WindowDecorator{window}
}
func (wd *WindowDecorator) Draw() {
	wd.window.Draw()
}
func (wd *WindowDecorator) GetDescription() string {
	return wd.window.GetDescription()
}
