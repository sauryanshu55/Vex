package main

import (
	"os"

	"github.com/openai/openai-go/v2"
)

func callAPI(userinput []string) {
	go buildMsgHist()                                                                       // Concurrently run buildMsgHist
	MsgTurn <- []openai.ChatCompletionMessageParamUnion{openai.SystemMessage(userinput[0])} // Send first command line arg to unblock buildMsgHist

	for { // Start channel, loop forever
		<-ApiTurn // Block API Call until request body with message history is formed

	}
}

func buildMsgHist() {
	sysIns := loadSysIns() // Load sysIns the first time, and let channel run

	// Build req the first time, reusable as long as channel is open
	var req = Request{
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{openai.SystemMessage(sysIns)}, // Add sysins as first, to message history
	}

	for { //Start channel, loop forever
		// Block execution here until MsgTurn channel recieves either userinput or assistant history
		hist := <-MsgTurn                            //  hist refers to both userinput and assistant history
		req.Messages = append(req.Messages, hist...) // Append to chat history
		ApiTurn <- req                               //Send request body to ApiTurn channel to unblock callAPI
	}

}

func loadSysIns() string {
	ins, err := os.ReadFile("SYS_INS.txt")
	if err != nil {
		return ""
	}
	return string(ins)
}
