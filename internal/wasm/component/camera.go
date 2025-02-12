package component

import (
	"syscall/js"
)

func CreateCamera() js.Value {

	h2 := js.Global().Get("document").Call("createElement", "h2")
	h2.Set("innerText", "Camera")

	return h2
}
