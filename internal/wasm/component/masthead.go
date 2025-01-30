package component

import (
	"syscall/js"
)

func CreateMasthead() js.Value {

	mastheadToggle := js.Global().Get("document").Call("createElement", "span")
	mastheadToggle.Set("className", "pf-v6-c-masthead__toggle")

	mastheadMain := js.Global().Get("document").Call("createElement", "div")
	mastheadMain.Set("className", "pf-v6-c-masthead__main")
	mastheadMain.Call("appendChild", mastheadToggle)

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead")
	masthead.Call("appendChild", mastheadMain)

	return masthead
}
