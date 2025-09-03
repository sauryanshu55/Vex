package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processInput(args []string) Response {
	return callAPI(args)
}

func processInitArg(initArgs string) {
	resp := processInput([]string{initArgs})

	if resp.Confidence {
		buildScript(resp.Commands)
		return
	}

	enterShell(resp)
}

func enterShell(initResp Response) {
	scanner := bufio.NewScanner(os.Stdin)
	currResp := initResp
	var commands []string

	for !currResp.Confidence {

		var args []string

		if currResp.NeedExplanation {
			fmt.Print(currResp.Question)
			if scanner.Scan() {
				arg := strings.TrimSpace(scanner.Text())
				args = append(args, arg)
			}
		}

		if currResp.NeedPrereqs {
			fmt.Print(currResp.Prereqs)
			if scanner.Scan() {
				arg := strings.TrimSpace(scanner.Text())
				args = append(args, arg)
			}
		}

		currResp := processInput(args)

		if !currResp.Confidence {
			commands = currResp.Commands
			break
		}

	}
	buildScript(commands)

}
