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

func (window Window) LocalStorage() LocalStorage {
	return LocalStorage{
		DOM: window.Get("localStorage"),
	}
}

func GetWindow() Window {
	return Window{
		DOM: GetGlobal(),
	}
}
