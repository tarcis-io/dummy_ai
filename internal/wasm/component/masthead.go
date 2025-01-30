package component

import (
	"syscall/js"
)

func CreateMasthead() js.Value {

	masthead := js.Global().Get("document").Call("createElement", "header")
	masthead.Set("className", "pf-v6-c-masthead")

	return masthead
}
