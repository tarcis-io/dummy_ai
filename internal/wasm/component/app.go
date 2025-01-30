package component

import (
	"syscall/js"
)

func CreateApp(pageContent js.Value) js.Value {

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", CreateMasthead())

	return page
}
