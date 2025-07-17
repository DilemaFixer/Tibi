package main

import (
	com	"github.com/DilemaFixer/Tibi/server/src/command"
)

func main() {
	_ , _ = com.ParseToCommand("comment add Started_working .some_var=true")
}
