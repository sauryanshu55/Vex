package main

import "github.com/openai/openai-go/v2"

// Its API's turn to run when this buffer recieves a complete request
var ApiTurn = make(chan Request)

// It's turn to build comprehensive message history when this buffer recives user request or an api response
var MsgTurn = make(chan []openai.ChatCompletionMessageParamUnion)
