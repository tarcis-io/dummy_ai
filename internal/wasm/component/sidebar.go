package component

import (
	"syscall/js"
)

func CreateSidebar() js.Value {

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Set("className", "pf-v6-c-page__sidebar")

	return sidebar
}
