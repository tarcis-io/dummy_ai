package util

const (
	app = iota
	appDescription
	mastheadNavigationButtonAriaLabel
	sidebarAboutNavigationItem
	sidebarHomeNavigationItem
)

var (
	allMessages = map[string]map[int]string{
		English: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navigation",
			sidebarAboutNavigationItem:        "About",
			sidebarHomeNavigationItem:         "Home",
		},
		Spanish: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegación",
			sidebarAboutNavigationItem:        "Sobre",
			sidebarHomeNavigationItem:         "Inicio",
		},
		Portuguese: {
			app:                               "DummyAI",
			appDescription:                    "Artificial intelligence for dummies",
			mastheadNavigationButtonAriaLabel: "Navegação",
			sidebarAboutNavigationItem:        "Sobre",
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

func SidebarAboutNavigationItem() string {

	return messages[sidebarAboutNavigationItem]
}

func SidebarHomeNavigationItem() string {

	return messages[sidebarHomeNavigationItem]
}
