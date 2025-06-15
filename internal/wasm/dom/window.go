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

func (window Window) Navigator() Navigator {

	return Navigator{
		DOM: window.Get("navigator"),
	}
}

func (window Window) OnLoad(listener func()) {

	window.AddEventListener("load", listener)
}

func GetWindow() Window {

	return Window{
		DOM: GetGlobal(),
	}
}
