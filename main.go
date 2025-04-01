package main

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	file := os.Args[1]
	fmt.Println(fmt.Sprintf("Compiling datapack from %s", file))

	state := lua.NewState()
	defer state.Close()

	if err := state.DoFile(file); err != nil {
		panic(err)
	}
}
