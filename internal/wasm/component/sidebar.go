package component

import (
	"syscall/js"
)

func CreateSidebar() js.Value {

	navigationList := js.Global().Get("document").Call("createElement", "ul")
	navigationList.Set("className", "pf-v6-c-nav__list")

	navigation := js.Global().Get("document").Call("createElement", "nav")
	navigation.Set("className", "pf-v6-c-nav")
	navigation.Call("appendChild", navigationList)

	sidebarBody := js.Global().Get("document").Call("createElement", "div")
	sidebarBody.Set("className", "pf-v6-c-page__sidebar-body")

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Set("className", "pf-v6-c-page__sidebar")
	sidebar.Call("appendChild", sidebarBody)

	return sidebar
}
