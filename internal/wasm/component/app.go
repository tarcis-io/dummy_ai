package component

import (
	"syscall/js"
)

func createPage() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("classList").Call("add", "pf-v6-c-page")

	return page
}
