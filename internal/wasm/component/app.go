package component

import (
	"syscall/js"
)

func createPage() js.Value {

	pageMainContainer := js.Global().Get("document").Call("createElement", "div")
	pageMainContainer.Get("classList").Call("add", "pf-v6-c-page__main-container")
	pageMainContainer.Set("tabindex", -1)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Get("classList").Call("add", "pf-v6-c-page")
	page.Call("appendChild", pageMainContainer)

	return page
}
