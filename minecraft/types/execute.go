package minecraft_types

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type Execute struct {
	Parameters []string
}

func RegisterExecute(L *lua.LState) {
	mt := L.NewTypeMetatable("Execute")
	L.SetGlobal("Execute", mt)
	L.SetField(mt, "new", L.NewFunction(newExecute))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), executeMethods))
}

func newExecute(L *lua.LState) int {
	execute := &Execute{}
	ud := L.NewUserData()
	ud.Value = execute
	L.SetMetatable(ud, L.GetTypeMetatable("Execute"))
	L.Push(ud)
	return 1
}

func checkExecute(L *lua.LState) *Execute {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Execute); ok {
		return v
	}
	L.ArgError(1, "Execute expected")
	return nil
}

var executeMethods = map[string]lua.LGFunction{
	"as":             executeAs,
	"ifScoreBetween": executeIfScoreBetween,
	"run":            executeRun,
}

// Implement methods

func executeAs(L *lua.LState) int {
	exec := checkExecute(L)
	exec.Parameters = append(exec.Parameters, fmt.Sprintf("as %s", L.CheckString(2)))
	// push self onto stack
	L.Push(L.CheckUserData(1))
	return 1
}

func executeIfScoreBetween(L *lua.LState) int {
	exec := checkExecute(L)
	num1 := -1
	num2 := -1
	comp := fmt.Sprintf("if score @s %s matches %d..%d", L.CheckString(2), num1, num2)
	exec.Parameters = append(exec.Parameters, comp)
	L.Push(L.CheckUserData(1))
	return 1
}

func executeRun(L *lua.LState) int {
	//TODO
	return 0
}
