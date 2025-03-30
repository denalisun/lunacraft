package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func mcSay(L *lua.LState) int {
	message := L.ToString(1)
	target := L.ToString(2)
	cmd := fmt.Sprintf("say %s %s", message, target)
	fmt.Println(cmd)
	return 0
}

func mcGive(L *lua.LState) int {
	item := L.ToString(1)
	target := L.ToString(2)
	cmd := fmt.Sprintf("give %s %s", target, item)
	fmt.Println(cmd)
	return 0
}

func main() {
	L := lua.NewState()
	defer L.Close()

	mc := L.NewTable()
	L.SetGlobal("mc", mc)

	L.SetField(mc, "say", L.NewFunction(mcSay))
	L.SetField(mc, "give", L.NewFunction(mcGive))

	// just testing for now
	script := `
	function tick()
		mc.say("hi", "@s")
		mc.give("minecraft:diamond_sword", "@s")
	end

	tick()
	`

	if err := L.DoString(script); err != nil {
		fmt.Println("lua err: ", err)
	}
}
