package util

const (
	app = iota
	appDescription
	mastheadNavigationButtonAriaLabel
	sidebarHomeNavigationItem
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navigation",
			sidebarHomeNavigationItem:         "Home",
		},
		Spanish: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegación",
			sidebarHomeNavigationItem:         "Inicio",
		},
		Portuguese: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegação",
			sidebarHomeNavigationItem:         "Início",
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

func MastheadNavigationButtonAriaLabel() string {

	return messages[mastheadNavigationButtonAriaLabel]
}

func SidebarHomeNavigationItem() string {

	return messages[sidebarHomeNavigationItem]
}
