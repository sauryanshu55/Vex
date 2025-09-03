package main

import (
	"fmt"
	"os"
	"runtime"
)

var SYS_OS string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vex \"<command>\"")
		os.Exit(1)
	}
	initArg := os.Args[1]

	SYS_OS = runtime.GOOS

	go buildMsgHist() // Start channel to build request

	processInitArg(initArg)

	close(ApiTurn)
	close(MsgTurn)
}
