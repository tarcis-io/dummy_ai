package component

import (
	"syscall/js"
)

func CreateApp() js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", CreateMasthead())
	page.Call("appendChild", CreateSidebar())

	return page
}
