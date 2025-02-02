package component

import (
	"syscall/js"
)

func CreateSidebar() js.Value {

	sidebarBody := js.Global().Get("document").Call("createElement", "div")
	sidebarBody.Set("className", "pf-v6-c-page__sidebar-body")

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Set("className", "pf-v6-c-page__sidebar")
	sidebar.Call("appendChild", sidebarBody)

	return sidebar
}
