package util

import (
	"syscall/js"
)

func IsDesktop() bool {

	breakpoint := js.Global().Call("getComputedStyle", js.Global().Get("document").Get("documentElement")).Call("getPropertyValue", "--pf-t--global--breakpoint--md").String()
	return js.Global().Call("matchMedia", "screen and (min-width: "+breakpoint+")").Get("matches").Bool()
}

func IsMobile() bool {

	return !IsDesktop()
}
