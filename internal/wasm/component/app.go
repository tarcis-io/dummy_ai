package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
)

func CreateApp(pageContent js.Value) js.Value {

	navigationIcon := js.Global().Get("document").Call("createElement", "i")
	navigationIcon.Set("className", "fas fa-bars")
	navigationIcon.Set("ariaHidden", true)

	navigationButtonIcon := js.Global().Get("document").Call("createElement", "span")
	navigationButtonIcon.Set("className", "pf-v6-c-button__icon")
	navigationButtonIcon.Call("appendChild", navigationIcon)

	navigationButton := js.Global().Get("document").Call("createElement", "button")
	navigationButton.Set("className", "pf-v6-c-button pf-m-plain")
	navigationButton.Set("type", "button")
	navigationButton.Set("ariaLabel", util.AppButtonNavigationAriaLabel())
	navigationButton.Call("appendChild", navigationButtonIcon)

	mastheadToggle := js.Global().Get("document").Call("createElement", "span")
	mastheadToggle.Set("className", "pf-v6-c-masthead__toggle")
	mastheadToggle.Call("appendChild", navigationButton)

	mastheadBrand := js.Global().Get("document").Call("createElement", "div")
	mastheadBrand.Set("className", "pf-v6-c-masthead__brand")

	mastheadMain := js.Global().Get("document").Call("createElement", "div")
	mastheadMain.Set("className", "pf-v6-c-masthead__main")
	mastheadMain.Call("appendChild", mastheadToggle)
	mastheadMain.Call("appendChild", mastheadBrand)

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead")
	masthead.Call("appendChild", mastheadMain)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", masthead)

	return page
}
