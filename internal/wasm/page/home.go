package main

import (
	"dummy_ai/internal/wasm/component"
)

func main() {

	component.CreateApp(component.CreateCamera())

	select {}
}
