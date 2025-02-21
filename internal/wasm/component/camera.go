package component

import (
	"syscall/js"
)

func CreateCamera() js.Value {

	return CreateCameraError("Title", "Text")
}

func CreateCameraError(title string, text string) js.Value {

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")

	return card
}
