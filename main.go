package main

import (
	"fmt"
	"runtime"

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
			fmt.Println(state)
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

	ExposableFunctions := []lua.RegistryFunction{{Name: "testFuncti", Function: binnenkaas}, {Name: "openWindow", Function: openGLFWindow}}

	lua.OpenLibraries(luaState)
	lua.SetFunctions(luaState, ExposableFunctions, 0)
	if err := lua.DoFile(luaState, "game.lua"); err != nil {
		panic(err)
	}

	// setupOpenGL.Startup()
}
