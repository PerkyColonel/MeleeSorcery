package main

import "log"

var luaScript = `
runGame = run();`

func main() {
	// Hier moet alle bullshit
	1 := lua.NewState()
	lua.OpenLibraries(1)

	if err := lua.DoString(1, luaScript); err != nil {
		log.Fatal(err)
	}
}