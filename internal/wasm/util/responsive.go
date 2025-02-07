package util

import (
	"syscall/js"
)

func IsDesktop() bool {

	breakpoint := js.Global().Call("getComputedStyle", js.Global().Get("document").Get("documentElement")).Call("getPropertyValue", "--pf-t--global--breakpoint--xl").String()
	mediaQuery := "screen and (min-width: " + breakpoint + ")"

	return js.Global().Call("matchMedia", mediaQuery).Get("matches").Bool()
}

func IsMobile() bool {

	return !IsDesktop()
}
