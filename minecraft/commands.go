package minecraft

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Tellraw(L *lua.LState) int {
	message := L.ToString(1)
	target := L.ToString(2)
	fmt.Printf("tellraw %s %s\n", target, message)
	return 0
}

func Give(L *lua.LState) int {
	id := L.ToString(1)
	count := L.ToInt(2)
	target := L.ToString(3)
	fmt.Printf("give %s %s %d\n", target, id, count)
	return 0
}

func Say(L *lua.LState) int {
	message := L.ToString(1)
	fmt.Printf("say %s\n", message)
	return 0
}

func RegisterCommands(L *lua.LState) {
	mcTable := L.NewTable()
	L.SetGlobal("mc", mcTable)

	L.SetField(mcTable, "tellraw", L.NewFunction(Tellraw))
	L.SetField(mcTable, "give", L.NewFunction(Give))
	L.SetField(mcTable, "say", L.NewFunction(Say))
}
