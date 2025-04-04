package minecraft

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// I gotta rewrite this
func Tellraw(L *lua.LState) int {
	message := L.ToString(1)
	target := L.ToString(2)
	cmd := fmt.Sprintf("tellraw %s \"%s\"", target, message)
	GlobalPackManager.FocusedFunction.Content = append(GlobalPackManager.FocusedFunction.Content, cmd)
	return 0
}

func Give(L *lua.LState) int {
	id := L.ToString(1)
	count := L.ToInt(2)
	target := L.ToString(3)
	cmd := fmt.Sprintf("give %s %s %d", target, id, count)
	GlobalPackManager.FocusedFunction.Content = append(GlobalPackManager.FocusedFunction.Content, cmd)
	return 0
}

func Say(L *lua.LState) int {
	message := L.ToString(1)
	cmd := fmt.Sprintf("say %s", message)
	GlobalPackManager.FocusedFunction.Content = append(GlobalPackManager.FocusedFunction.Content, cmd)
	return 0
}

func RunCommand(L *lua.LState) int {
	cmd := L.ToString(1)
	GlobalPackManager.FocusedFunction.Content = append(GlobalPackManager.FocusedFunction.Content, cmd)

	return 0
}

func RegisterCommands(L *lua.LState) {
	mcTable := L.NewTable()
	L.SetGlobal("mc", mcTable)
	L.SetField(mcTable, "tellraw", L.NewFunction(Tellraw))
	L.SetField(mcTable, "give", L.NewFunction(Give))
	L.SetField(mcTable, "say", L.NewFunction(Say))
	L.SetField(mcTable, "runCommand", L.NewFunction(RunCommand))
}
