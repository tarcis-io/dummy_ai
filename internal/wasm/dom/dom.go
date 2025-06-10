package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func GetGlobal() DOM {

	return DOM{
		jsObject: js.Global(),
	}
}
