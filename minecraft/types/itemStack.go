package minecraft_types

import (
	lua "github.com/yuin/gopher-lua"
)

type ItemStack struct {
	Id         string
	Components map[string]any
}

func RegisterItemStack(L *lua.LState) {
	mt := L.NewTypeMetatable("ItemStack")
	L.SetGlobal("ItemStack", mt)
	L.SetField(mt, "new", L.NewFunction(newItemStack))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ItemStackMethods))
}

func newItemStack(L *lua.LState) int {
	stack := &ItemStack{L.ToString(1), map[string]any{}}
	ud := L.NewUserData()
	ud.Value = stack
	L.SetMetatable(ud, L.GetTypeMetatable("ItemStack"))
	L.Push(ud)
	return 1
}

func checkItemStack(L *lua.LState) *ItemStack {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*ItemStack); ok {
		return v
	}
	L.ArgError(1, "ItemStack expected")
	return nil
}

var ItemStackMethods = map[string]lua.LGFunction{}

func AddNewComponent(L *lua.LState) int {

	return 0
}
