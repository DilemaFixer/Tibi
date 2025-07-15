package main

import (
	com	"github.com/DilemaFixer/Tibi/server/src/command"
)

func main() {
	_ , _ = com.parseToCommand("tibi comment add 15 Started working .some_var=true")
}
