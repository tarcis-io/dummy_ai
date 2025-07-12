package dom

type (
	Window struct {
		DOM
	}
)

func (window Window) Document() Document {
	return Document{
		DOM: window.Get("document"),
	}
}

func GetWindow() Window {
	return Window{
		DOM: GetGlobal(),
	}
}
