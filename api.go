package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func callAPI(args []string) Response {
	// FUNC OVERVIEW:
	// Send args to MsgTurn channel
	// Block here until ApiTurn recieves anything back
	// Requst recieved
	// Process it
	// Send assistant response to MsgTurn channel
	msgs := make([]openai.ChatCompletionMessageParamUnion, 0, len(args)) // Make a list of openai user msg type
	for _, s := range args {
		msgs = append(msgs, openai.SystemMessage(s)) // Convert and populate array of openai msg type
	}
	MsgTurn <- msgs  // Send to MsgTurn channel
	req := <-ApiTurn // Block execution here until proper req is returned
	// Set up API Client
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
	} // Set up API params
	rawResp, err := client.Chat.Completions.New(ctx, params) // Get raw response
	if err != nil {
	}
	resp := parseResp(rawResp.Choices[0].Message.Content) // Parse it

	var asstMsgs []openai.ChatCompletionMessageParamUnion // Make a list of openai asst msg type
	if resp.NeedExplanation {
		asstMsgs = append(asstMsgs, openai.AssistantMessage(resp.Question)) // populate as msg history
	}
	if resp.NeedPrereqs {
		asstMsgs = append(asstMsgs, openai.AssistantMessage(resp.Prereqs))
	}

	MsgTurn <- asstMsgs // Send to MsgTurn so msg history can be built

	return resp
}

func buildMsgHist() {
	sysIns := loadSysIns() // Load sys ins from file

	// Intialize a req body. Keep channel open concurrently so that a new req body doesn't need to be created everytime
	var req = Request{
		Seed:     openai.Int(0),
		Model:    openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{openai.SystemMessage(sysIns)}, // Add sysins as a first, to message history
	}

	// Loop this part forever
	//		Block here until MsgTurn recieves either userinput, or assistant response
	// 		Process the userinput/assistant response
	// 		Send request with complete history ApiTurn channel
	// 		(This loops, but the next line will block until a new userinput or an assistant response is recieved)

	for { // Loop indefinitely
		hist := <-MsgTurn                            // Listen for any asst msg or user msg
		req.Messages = append(req.Messages, hist...) // Append it to req body to build complete history
		ApiTurn <- req                               // Send it to channel so req is ready to be used
	}
}
