package minecraft_types

import (
	"github.com/denalisun/lunacraft/minecraft"
	lua "github.com/yuin/gopher-lua"
)

func RegisterLunaType(L *lua.LState) {
	table := L.NewTable()
	L.SetGlobal("luna", table)
	L.SetField(table, "tick", L.NewFunction(lunaTick))
	L.SetField(table, "load", L.NewFunction(lunaLoad))
	L.SetField(table, "declare", L.NewFunction(lunaDeclare))
}

func lunaLoad(L *lua.LState) int {
	callback := L.CheckFunction(1)
	fnc := minecraft.GlobalPackManager.AddNewFunction("load")
	currentFocused := minecraft.GlobalPackManager.FocusedFunction
	minecraft.GlobalPackManager.FocusedFunction = fnc
	err := L.CallByParam(lua.P{
		Fn:      callback,
		NRet:    0,
		Protect: true,
	})
	if err != nil {
		panic(err)
	}
	minecraft.GlobalPackManager.FocusedFunction = currentFocused
	return 0
}

func lunaTick(L *lua.LState) int {
	callback := L.CheckFunction(1)
	fnc := minecraft.GlobalPackManager.AddNewFunction("tick")
	currentFocused := minecraft.GlobalPackManager.FocusedFunction
	minecraft.GlobalPackManager.FocusedFunction = fnc
	err := L.CallByParam(lua.P{
		Fn:      callback,
		NRet:    0,
		Protect: true,
	})
	if err != nil {
		panic(err)
	}
	minecraft.GlobalPackManager.FocusedFunction = currentFocused
	return 0
}

// worst code i've ever written but okay
func lunaDeclare(L *lua.LState) int {
	callback := L.CheckFunction(2)
	fnc := minecraft.GlobalPackManager.AddNewFunction(L.CheckString(1))
	currentFocused := minecraft.GlobalPackManager.FocusedFunction
	minecraft.GlobalPackManager.FocusedFunction = fnc
	err := L.CallByParam(lua.P{
		Fn:      callback,
		NRet:    0,
		Protect: true,
	})
	if err != nil {
		panic(err)
	}
	minecraft.GlobalPackManager.FocusedFunction = currentFocused
	return 0
}
