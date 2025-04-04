package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/denalisun/lunacraft/minecraft"
	minecraft_types "github.com/denalisun/lunacraft/minecraft/types"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	file := os.Args[1]
	fmt.Printf("Compiling datapack from %s\n", file)

	state := lua.NewState()
	defer state.Close()

	minecraft_types.RegisterLunaType(state)
	minecraft_types.RegisterPlayerSelector(state)
	minecraft_types.RegisterExecute(state)
	minecraft_types.RegisterItemStack(state)
	minecraft.RegisterCommands(state)

	if err := state.DoFile(file); err != nil {
		panic(err)
	}

	for i := 0; i < len(minecraft.GlobalPackManager.Functions); i++ {
		fnc := minecraft.GlobalPackManager.Functions[i]
		fmt.Printf("Compiling %s\n", fnc.Name)
		joined := strings.Join(fnc.Content, "\n")
		bData := []byte(joined)
		if err := os.WriteFile(fmt.Sprintf("%s.mcfunction", fnc.Name), bData, 0644); err != nil {
			log.Fatal(err)
		}
	}
}
