package component

import (
	"syscall/js"
)

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	return page
}
