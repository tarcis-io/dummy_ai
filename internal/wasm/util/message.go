package util

const (
	app = iota
	appDescription
	ariaLabelNavigation
	navigationItemAbout
	navigationItemHome
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navigation",
			navigationItemAbout: "About",
			navigationItemHome:  "Home",
		},
		Spanish: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navigation",
			navigationItemAbout: "Sobre",
			navigationItemHome:  "Inicio",
		},
		Portuguese: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			ariaLabelNavigation: "Navegação",
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
