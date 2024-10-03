package main

import (
	"runtime"

	"github.com/PerkyColonel/MeleeSorcery/tree/main/drawing"
	setupOpenGL "github.com/PerkyColonel/MeleeSorcery/tree/main/window"
	"github.com/Shopify/go-lua"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}
func main() {

	binnenkaas := lua.Function(
		func(state *lua.State) int {
			// state.Field(0, "Waa")

			return 0
		},
	)

	drawQuadrino := lua.Function(
		func(state *lua.State) int {
			x, _ := state.ToNumber(1)
			y, _ := state.ToNumber(2)
			w, _ := state.ToNumber(3)
			h, _ := state.ToNumber(4)

			drawing.AddQuad(x, y, w, h)
			return 0
		},
	)

	openGLFWindow := lua.Function(
		func(state *lua.State) int {
			setupOpenGL.Startup()
			return 0
		},
	)

	luaState := lua.NewState()
	lua.BaseOpen(luaState)

	ExposableFunctions := []lua.RegistryFunction{{Name: "testFuncti", Function: binnenkaas}, {Name: "openWindow", Function: openGLFWindow}, {Name: "addQuad", Function: drawQuadrino}}

	lua.OpenLibraries(luaState)
	lua.SetFunctions(luaState, ExposableFunctions, 0)
	if err := lua.DoFile(luaState, "game.lua"); err != nil {
		panic(err)
	}

	// setupOpenGL.Startup()
}
