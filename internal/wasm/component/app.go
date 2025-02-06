package component

import (
	"syscall/js"
)

func CreateApp(pageContent js.Value) js.Value {

	pageMainContainer := js.Global().Get("document").Call("createElement", "div")
	pageMainContainer.Set("className", "pf-v6-c-page__main-container pf-m-fill")
	pageMainContainer.Set("tabIndex", -1)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Get("style").Call("setProperty", "--pf-v6-c-page__sidebar--xl--TranslateX", "-100%")
	page.Call("appendChild", CreateMasthead())
	page.Call("appendChild", CreateSidebar())
	page.Call("appendChild", pageMainContainer)

	return page
}
