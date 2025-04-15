package dom

type (
	Window struct {
		DOM
	}
)

func (window Window) Document() Document {

	return Document{
		DOM: DOM{
			jsObject: window.jsObject.Get("document"),
		},
	}
}

func GetWindow() Window {

	return Window{
		DOM: DOM{
			jsObject: GetGlobal(),
		},
	}
}
