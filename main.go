package main

import (
	"fmt"
	"runtime"

	setupOpenGL "github.com/PerkyColonel/MeleeSorcery/tree/main/window"
	"github.com/Shopify/go-lua"
	"github.com/go-gl/glfw/v3.3/glfw"
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
			err := glfw.Init()
			if err != nil {
				panic(err)
			}
			defer glfw.Terminate()

			window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
			if err != nil {
				panic(err)
			}

			window.MakeContextCurrent()
			windowLoop(window)
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

	setupOpenGL.Startup()
}

func windowLoop(window *glfw.Window) {
	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
