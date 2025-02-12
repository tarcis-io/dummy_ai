package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
)

func CreateApp(pageContent js.Value) {

	html := js.Global().Get("document").Get("documentElement")
	html.Set("lang", util.Language())

	pageMainBody := js.Global().Get("document").Call("createElement", "div")
	pageMainBody.Set("className", "pf-v6-c-page__main-body")
	pageMainBody.Call("appendChild", pageContent)

	pageMainSection := js.Global().Get("document").Call("createElement", "section")
	pageMainSection.Set("className", "pf-v6-c-page__main-section pf-m-fill")
	pageMainSection.Call("appendChild", pageMainBody)

	pageMain := js.Global().Get("document").Call("createElement", "main")
	pageMain.Set("className", "pf-v6-c-page__main")
	pageMain.Set("tabIndex", -1)
	pageMain.Call("appendChild", pageMainSection)

	pageMainContainer := js.Global().Get("document").Call("createElement", "div")
	pageMainContainer.Set("className", "pf-v6-c-page__main-container pf-m-fill")
	pageMainContainer.Set("tabIndex", -1)
	pageMainContainer.Call("appendChild", pageMain)

	page := js.Global().Get("document").Call("createElement", "div")
	page.Set("className", "pf-v6-c-page")
	page.Get("style").Call("setProperty", "--pf-v6-c-page__sidebar--xl--TranslateX", "-100%")
	page.Call("appendChild", CreateMasthead())
	page.Call("appendChild", CreateSidebar())
	page.Call("appendChild", pageMainContainer)

	body := js.Global().Get("document").Get("body")
	body.Call("appendChild", page)
}
