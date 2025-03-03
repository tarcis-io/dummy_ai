package util

const (
	app = iota
	appDescription
	appDevelopedBy
	appVersion
	ariaLabelNavigation
	buttonReloadPage
	buttonReturnToHome
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
			appDevelopedBy:      "Developed by t@rcis.io",
			appVersion:          "Version 0.0.1",
			ariaLabelNavigation: "Navigation",
			buttonReloadPage:    "Reload page",
			buttonReturnToHome:  "Return to home",
			error404:            "Error 404",
			error404Message:     "That page no longer exists",
			navigationItemAbout: "About",
			navigationItemHome:  "Home",
		},
		Spanish: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			appDevelopedBy:      "Desarrollado por t@rcis.io",
			appVersion:          "Versión 0.0.1",
			ariaLabelNavigation: "Navegación",
			buttonReloadPage:    "Recargar página",
			buttonReturnToHome:  "Volver al inicio",
			error404:            "Error 404",
			error404Message:     "Esa página ya no existe",
			navigationItemAbout: "Sobre",
			navigationItemHome:  "Inicio",
		},
		Portuguese: {
			app:                 "DummyAI",
			appDescription:      "Artificial intelligence for dummies",
			appDevelopedBy:      "Desenvolvido por t@rcis.io",
			appVersion:          "Versão 0.0.1",
			ariaLabelNavigation: "Navegação",
			buttonReloadPage:    "Recarregar página",
			buttonReturnToHome:  "Voltar para o início",
			error404:            "Erro 404",
			error404Message:     "Essa página não existe mais",
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

func AppDevelopedBy() string {

	return messages[appDevelopedBy]
}

func AppVersion() string {

	return messages[appVersion]
}

func AriaLabelNavigation() string {

	return messages[ariaLabelNavigation]
}

func ButtonReloadPage() string {

	return messages[buttonReloadPage]
}

func ButtonReturnToHome() string {

	return messages[buttonReturnToHome]
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
