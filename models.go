package main

import (
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/packages/param"
	"github.com/openai/openai-go/v2/shared"
)

type Response struct {
	NeedExplanation bool     `json:"need_explanation"`
	Question        string   `json:"question"`
	NeedPrereqs     bool     `json:"need_prereqs"`
	Prereqs         string   `json:"prereqs"`
	Risky           bool     `json:"risky"`
	Commands        []string `json:"commands"`
	Confidence      bool     `json:"confidence"`
}

type Request struct {
	Messages []openai.ChatCompletionMessageParamUnion
	Seed     param.Opt[int64]
	Model    shared.ChatModel
}
