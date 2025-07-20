package main

import (
	"fmt"
	"log"

	com "github.com/DilemaFixer/Tibi/server/src/command"
)

func main() {
	log.SetFlags(log.Lmsgprefix)
	log.SetPrefix("[SERVER] ")

	str := "comment add \"Started_working\" .some_var=true .some_varr=12 --some_flag"
	command, err := com.ParseToCommand(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(command)
	}
	r := com.NewRouter(errorHandler)
	r.NewEndPointGroup("comment").AddEndPoint("add", test).EndFilling().EndFilling()
	r.Route(*command)
}

func errorHandler(err error) {
	fmt.Println(err)
}

func test(signs []com.Significance) error {
	for _, sign := range signs {
		fmt.Println(sign)
	}
	return nil
}
