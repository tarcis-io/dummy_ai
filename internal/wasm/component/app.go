package component

import (
	"syscall/js"
)

func CreateApp(pageContent js.Value) js.Value {

	mastheadMain := js.Global().Get("document").Call("createElement", "div")
	mastheadMain.Set("classNAme", "pf-v6-c-masthead__main")

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead")
	masthead.Call("appendChild", mastheadMain)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", masthead)

	return page
}
