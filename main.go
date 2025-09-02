package main

import (
	"fmt"
	"os"
)

func processInput(userinput string) {
	go callAPI([]string{userinput})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vex \"<command>\"")
		os.Exit(1)
	}
	userinput := os.Args[1]

	processInput(userinput)

}
