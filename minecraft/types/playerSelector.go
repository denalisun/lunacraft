package minecraft_types

import lua "github.com/yuin/gopher-lua"

func RegisterPlayerSelector(L *lua.LState) {
	table := L.NewTable()
	L.SetField(table, "ALL", lua.LString("@a"))
	L.SetField(table, "SELF", lua.LString("@s"))
	L.SetGlobal("Selector", table)
}
