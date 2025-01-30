package component

import (
	"syscall/js"
)

func CreateMasthead() js.Value {

	navigationIcon := js.Global().Get("document").Call("createElement", "i")
	navigationIcon.Set("className", "fas fa-bars")
	navigationIcon.Set("ariaHidden", true)

	navigationButtonIcon := js.Global().Get("document").Call("createElement", "span")
	navigationButtonIcon.Set("className", "pf-v6-c-button__icon")
	navigationButtonIcon.Call("appendChild", navigationIcon)

	navigationButton := js.Global().Get("document").Call("createElement", "button")
	navigationButton.Set("className", "pf-v6-c-button pf-m-plain")
	navigationButton.Call("appendChild", navigationButtonIcon)

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
