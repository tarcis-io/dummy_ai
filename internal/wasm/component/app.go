package component

import (
	"syscall/js"
)

func createAppPage(pageContent js.Value) js.Value {

	pageMainContainer := js.Global().Get("document").Call("createElement", "div")
	pageMainContainer.Set("className", "pf-v6-page__main-container")
	pageMainContainer.Set("tabindex", -1)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", pageMainContainer)

	return page
}
