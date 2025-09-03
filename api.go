package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func callAPI(userinput []string) {
	// Send userinput to MsgTurn channel
	// Block here ApiTurn recieves anything back
	// Requst recieved
	// Process it
	// Send assistant response to MsgTurn channel

	messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(userinput))
	for _, s := range userinput {
		messages = append(messages, openai.SystemMessage(s))
	}
	MsgTurn <- messages
	req := <-ApiTurn

	err := godotenv.Load()
	if err != nil {
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	ctx := context.Background()

	params := openai.ChatCompletionNewParams{
		Messages: req.Messages,
		Seed:     req.Seed,
		Model:    req.Model,
	}
	resp, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
	}
}

func buildMsgHist() {
	// Intialize a request body
	sysIns := loadSysIns()

	var req = Request{
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{openai.SystemMessage(sysIns)}, // Add sysins as first, to message history
	}

	// Loop this part forever
	//		Block here until MsgTurn recieves either userinput, or assistant response
	// 		Process the userinput/assistant response
	// 		Send request with complete history ApiTurn channel
	// 		(This loops, but the next line will block until a new userinput or an assistant response is recieved)

	for {
		hist := <-MsgTurn
		req.Messages = append(req.Messages, hist...)
		ApiTurn <- req
	}
}

func loadSysIns() string {
	ins, err := os.ReadFile("SYS_INS.txt")
	if err != nil {
		return ""
	}
	return string(ins)
}
