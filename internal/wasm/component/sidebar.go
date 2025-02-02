package component

import (
	"syscall/js"
)

import (
	"dummy_ai/internal/wasm/util"
)

func CreateSidebar() js.Value {

	navigationItems := map[string]string{
		util.SidebarHomeNavigationItem():  "/",
		util.SidebarAboutNavigationItem(): "/about",
	}

	navigationList := js.Global().Get("document").Call("createElement", "ul")
	navigationList.Set("className", "pf-v6-c-nav__list")
	navigationList.Set("role", "list")

	for text, href := range navigationItems {

		navigationList.Call("appendChild", createSidebarNavigationItem(text, href))
	}

	navigation := js.Global().Get("document").Call("createElement", "nav")
	navigation.Set("className", "pf-v6-c-nav")
	navigation.Call("appendChild", navigationList)

	sidebarBody := js.Global().Get("document").Call("createElement", "div")
	sidebarBody.Set("className", "pf-v6-c-page__sidebar-body")
	sidebarBody.Call("appendChild", navigation)

	sidebar := js.Global().Get("document").Call("createElement", "div")
	sidebar.Set("className", "pf-v6-c-page__sidebar")
	sidebar.Call("appendChild", sidebarBody)

	return sidebar
}

func createSidebarNavigationItem(text string, href string) js.Value {

	navigationLinkText := js.Global().Get("document").Call("createElement", "span")
	navigationLinkText.Set("className", "pf-v6-c-nav__link-text")
	navigationLinkText.Set("innerText", text)

	navigationLink := js.Global().Get("document").Call("createElement", "a")
	navigationLink.Set("className", "pf-v6-c-nav__link")
	navigationLink.Set("href", href)
	navigationLink.Call("appendChild", navigationLinkText)

	if pathname := js.Global().Get("location").Get("pathname").String(); pathname == href {

		navigationLink.Set("ariaCurrent", "page")
		navigationLink.Get("classList").Call("add", "pf-m-current")
	}

	navigationItem := js.Global().Get("document").Call("createElement", "li")
	navigationItem.Set("className", "pf-v6-c-nav__item")
	navigationItem.Call("appendChild", navigationLink)

	return navigationItem
}
