package main

import (
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/packages/param"
	"github.com/openai/openai-go/v2/shared"
)

type Request struct {
	Messages []openai.ChatCompletionMessageParamUnion
	Seed     param.Opt[int64]
	Model    shared.ChatModel
}

type Response struct {
	NeedExplanation bool     `json:"need_explanation"`
	Question        string   `json:"question"`
	NeedPrereqs     bool     `json:"need_prereqs"`
	Prereqs         string   `json:"prereqs"`
	Risky           bool     `json:"risky"`
	Commands        []string `json:"commands"`
	Confidence      bool     `json:"confidence"`
}

type RawResponse struct {
	NeedExplanation string   `json:"need_explanation"`
	Question        string   `json:"question"`
	NeedPrereqs     string   `json:"need_prereqs"`
	Prereqs         string   `json:"prereqs"`
	Risky           string   `json:"risky"`
	Commands        []string `json:"commands"`
	Confidence      string   `json:"confidence"`
}
