package main

import (
	com "github.com/DilemaFixer/Tibi/server/src/command"
	"log"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[SERVER] ")

	str := "comment add \"Started_working\" .some_var=true --some_flag"
	_ , _ = com.ParseToCommand(str)
}


