package main

import (
	com "github.com/DilemaFixer/Tibi/server/src/command"
	"log"
	"fmt"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[SERVER] ")

	str := "comment add \"Started_working\" .some_var=true --some_flag"
	command , err := com.ParseToCommand(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(command)
	}
}


