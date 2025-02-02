package component

import (
	"syscall/js"
)

func CreateSidebar() js.Value {

	navigationList := js.Global().Get("document").Call("createElement", "ul")
	navigationList.Set("className", "pf-v6-c-nav__list")
	navigationList.Set("role", "list")

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
	navigationLink.Call("appendChild", navigationLink)

	navigationItem := js.Global().Get("document").Call("createElement", "li")
	navigationItem.Set("className", "pf-v6-c-nav__item")
	navigationItem.Call("appendChild", navigationLink)

	return navigationItem
}
