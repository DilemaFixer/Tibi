package main

import (
	"fmt"
	"log"

	com "github.com/DilemaFixer/Tibi/server/src/command"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[SERVER] ")

	str := "comment add \"Started_working\" .some_var=true --some_flag"
	command, err := com.ParseToCommand(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(command)
	}
	var r *com.Router
	r.NewEndPointGroup("test").AddEndPoint("str", nil).EndFilling().EndFilling().NewEndPointGroup("test2")
}
