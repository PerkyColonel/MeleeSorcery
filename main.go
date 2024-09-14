package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/Shopify/go-lua"
)

func main() {

	binnenkaas := lua.Function(
		func(state *lua.State) int {
			fmt.Println(state)
			return 0
		},
	)

	openWindow := lua.Function(
		func(state *lua.State) int {
			a := app.New()
			w := a.NewWindow("Hello")

			hello := widget.NewLabel("Hello Fyne!")
			w.SetContent(container.NewVBox(
				hello,
				widget.NewButton("Hi!", func() {
					hello.SetText("Welcome :)")
				}),
			))

			w.ShowAndRun()
			return 0
		},
	)

	luaState := lua.NewState()
	lua.BaseOpen(luaState)

	ExposableFunctions := []lua.RegistryFunction{{Name: "testFuncti", Function: binnenkaas}, {Name: "openWindow", Function: openWindow}}

	lua.OpenLibraries(luaState)
	lua.SetFunctions(luaState, ExposableFunctions, 0)
	if err := lua.DoFile(luaState, "game.lua"); err != nil {
		panic(err)
	}
}
