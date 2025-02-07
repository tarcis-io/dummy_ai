package util

const (
	app = iota
	appDescription
	ariaLabelNavigation
	error404
	error404Message
	navigationItemAbout
	navigationItemHome
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navigation",
			error404:            "Error 404",
			error404Message:     "",
			navigationItemAbout: "About",
			navigationItemHome:  "Home",
		},
		Spanish: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navegación",
			error404:            "Error 404",
			error404Message:     "",
			navigationItemAbout: "Sobre",
			navigationItemHome:  "Inicio",
		},
		Portuguese: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navegação",
			error404:            "Error 404",
			error404Message:     "",
			navigationItemAbout: "Sobre",
			navigationItemHome:  "Início",
		},
	}
	messages = allMessages[language]
)

func App() string {

	return messages[app]
}

func AppDescription() string {

	return messages[appDescription]
}

func AriaLabelNavigation() string {

	return messages[ariaLabelNavigation]
}

func Error404() string {

	return messages[error404]
}

func Error404Message() string {

	return messages[error404Message]
}

func NavigationItemAbout() string {

	return messages[navigationItemAbout]
}

func NavigationItemHome() string {

	return messages[navigationItemHome]
}
