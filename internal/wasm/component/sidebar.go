package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
)

func init() {

	js.Global().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) any {

		sidebarClassList := js.Global().Get("document").Call("getElementById", "sidebar").Get("classList")

		if util.IsDesktop() {

			sidebarClassList.Call("add", "pf-m-expanded")
			return nil
		}

		sidebarClassList.Call("remove", "pf-m-expanded")
		return nil
	}))
}

func CreateSidebar() js.Value {

	navigationItems := [][]string{
		{
			"/",
			util.NavigationItemHome(),
		},
		{
			"/about",
			util.NavigationItemAbout(),
		},
	}

	navigationList := js.Global().Get("document").Call("createElement", "ul")
	navigationList.Set("className", "pf-v6-c-nav__list")
	navigationList.Set("role", "list")

	for _, navigationItem := range navigationItems {

		navigationList.Call("appendChild", CreateSidebarNavigationItem(navigationItem[0], navigationItem[1]))
	}

	navigation := js.Global().Get("document").Call("createElement", "nav")
	navigation.Set("className", "pf-v6-c-nav")
	navigation.Call("appendChild", navigationList)

	sidebarBody := js.Global().Get("document").Call("createElement", "div")
	sidebarBody.Set("className", "pf-v6-c-page__sidebar-body")
	sidebarBody.Call("appendChild", navigation)

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Set("className", "pf-v6-c-page__sidebar")
	sidebar.Set("id", "sidebar")
	sidebar.Call("appendChild", sidebarBody)

	if util.IsDesktop() {

		sidebar.Get("classList").Call("add", "pf-m-expanded")
	}

	return sidebar
}

func CreateSidebarNavigationItem(href string, text string) js.Value {

	navigationLinkText := js.Global().Get("document").Call("createElement", "span")
	navigationLinkText.Set("className", "pf-v6-c-nav__link-text")
	navigationLinkText.Set("innerText", text)

	navigationLink := js.Global().Get("document").Call("createElement", "a")
	navigationLink.Set("className", "pf-v6-c-nav__link")
	navigationLink.Set("href", href)
	navigationLink.Call("appendChild", navigationLinkText)

	if pathname := js.Global().Get("location").Get("pathname").String(); pathname == href {

		navigationLink.Get("classList").Call("add", "pf-m-current")
		navigationLink.Set("ariaCurrent", "page")
	}

	navigationItem := js.Global().Get("document").Call("createElement", "li")
	navigationItem.Set("className", "pf-v6-c-nav__item")
	navigationItem.Call("appendChild", navigationLink)

	return navigationItem
}

func ToggleSidebar() {

	js.Global().Get("document").Call("getElementById", "sidebar").Get("classList").Call("toggle", "pf-m-expanded")
}
