package component

import (
	"syscall/js"
)

func CreateApp(pageContent js.Value) js.Value {

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead")

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", masthead)

	return page
}
