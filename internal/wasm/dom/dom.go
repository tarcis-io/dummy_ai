package dom

import (
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func Global() DOM {

	return DOM{
		jsObject: js.Global(),
	}
}
