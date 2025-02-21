package component

import (
	"syscall/js"
)

func CreateCamera() js.Value {

	return CreateCameraError("Title", "Text")
}

func CreateCameraError(title string, text string) js.Value {

	cardTitle := js.Global().Get("document").Call("createElement", "div")
	cardTitle.Set("className", "pf-v6-c-card__title")

	card := js.Global().Get("document").Call("createElement", "div")
	card.Set("className", "pf-v6-c-card")
	card.Call("appendChild", cardTitle)

	return card
}
