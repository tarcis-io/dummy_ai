package component

import (
	"syscall/js"
)

func createAppPage(pageContent js.Value) js.Value {

	pageMainSection := js.Global().Get("document").Call("createElement", "section")
	pageMainSection.Set("className", "pf-v6-c-page__main-section")

	pageMain := js.Global().Get("document").Call("createElement", "main")
	pageMain.Set("className", "pf-v6-c-page__main")
	pageMain.Set("tabindex", -1)
	pageMain.Call("appendChild", pageMainSection)

	pageMainContainer := js.Global().Get("document").Call("createElement", "div")
	pageMainContainer.Set("className", "pf-v6-c-page__main-container")
	pageMainContainer.Set("tabindex", -1)
	pageMainContainer.Call("appendChild", pageMain)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Call("appendChild", pageMainContainer)

	return page
}
