package main

import (
	"fmt"
	"os"

	"github.com/denalisun/lunacraft/minecraft"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	file := os.Args[1]
	fmt.Printf("Compiling datapack from %s\n", file)

	state := lua.NewState()
	defer state.Close()

	minecraft.RegisterCommands(state)

	if err := state.DoFile(file); err != nil {
		panic(err)
	}
}
