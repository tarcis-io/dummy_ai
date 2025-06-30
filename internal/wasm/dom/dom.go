package dom

import (
	"errors"
	"syscall/js"
)

type (
	DOM struct {
		jsObject js.Value
	}
)

func (dom DOM) Bool() bool {

	return dom.jsObject.Bool()
}

func (dom DOM) String() string {

	return dom.jsObject.String()
}

func (dom DOM) Int() int {

	return dom.jsObject.Int()
}

func (dom DOM) Float() float64 {

	return dom.jsObject.Float()
}

func (dom DOM) Truthy() bool {

	return dom.jsObject.Truthy()
}

func (dom DOM) Get(property string) DOM {

	return DOM{
		jsObject: dom.jsObject.Get(property),
	}
}

func (dom DOM) Set(property string, value any) {

	dom.jsObject.Set(property, value)
}

func (dom DOM) Call(method string, arguments ...any) DOM {

	for index, argument := range arguments {

		switch argument := argument.(type) {

		case DOM:

			arguments[index] = argument.jsObject

		}
	}

	return DOM{
		jsObject: dom.jsObject.Call(method, arguments...),
	}
}

func (dom DOM) Await() (DOM, error) {

	value := make(chan DOM)
	defer close(value)

	onResolve := js.FuncOf(func(this js.Value, arguments []js.Value) any {

		value <- DOM{
			jsObject: arguments[0],
		}

		return nil
	})
	defer onResolve.Release()

	err := make(chan error)
	defer close(err)

	onReject := js.FuncOf(func(this js.Value, arguments []js.Value) any {

		err <- errors.New(arguments[0].String())
		return nil
	})
	defer onReject.Release()

	dom.Call("then", onResolve, onReject)

	select {

	case value := <-value:

		return value, nil

	case err := <-err:

		return DOM{}, err

	}
}

func GetGlobal() DOM {

	return DOM{
		jsObject: js.Global(),
	}
}
