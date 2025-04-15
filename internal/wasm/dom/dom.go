package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func GetGlobal() js.Value {

	return js.Global()
}
