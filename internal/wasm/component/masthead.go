package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
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
	navigationButton.Set("type", "button")
	navigationButton.Set("ariaLabel", util.MastheadNavigationButtonAriaLabel())
	navigationButton.Call("appendChild", navigationButtonIcon)

	mastheadToggle := js.Global().Get("document").Call("createElement", "span")
	mastheadToggle.Set("className", "pf-v6-c-masthead__toggle")
	mastheadToggle.Call("appendChild", navigationButton)

	img := js.Global().Get("document").Call("createElement", "img")
	img.Set("src", "/logo.svg")
	img.Set("alt", "")

	title := js.Global().Get("document").Call("createElement", "div")
	title.Set("className", "pf-v6-c-title pf-m-h1 pf-m-page-title pf-v6-u-display-inline pf-v6-u-ml-md pf-v6-u-text-color-regular")
	title.Set("innerText", util.App())

	mastheadLogo := js.Global().Get("document").Call("createElement", "a")
	mastheadLogo.Set("className", "pf-v6-c-masthead__logo")
	mastheadLogo.Set("href", "/")
	mastheadLogo.Get("style").Set("textDecoration", "none")
	mastheadLogo.Call("appendChild", img)
	mastheadLogo.Call("appendChild", title)

	mastheadBrand := js.Global().Get("document").Call("createElement", "div")
	mastheadBrand.Set("className", "pf-v6-c-masthead__brand")
	mastheadBrand.Call("appendChild", mastheadLogo)

	mastheadMain := js.Global().Get("document").Call("createElement", "div")
	mastheadMain.Set("className", "pf-v6-c-masthead__main")
	mastheadMain.Call("appendChild", mastheadToggle)
	mastheadMain.Call("appendChild", mastheadBrand)

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead pf-m-display-inline")
	masthead.Call("appendChild", mastheadMain)

	return masthead
}
