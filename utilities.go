package main

import (
	"encoding/json"
	"os"
	"strings"
)

func loadSysIns() string {
	ins, err := os.ReadFile("SYS_INS.txt")
	if err != nil {
		return ""
	}
	return string(ins)
}

func parseResp(input string) Response {
	clean := strings.TrimSpace(input)
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimSuffix(clean, "```")
	clean = strings.TrimSpace(clean)

	var raw RawResponse
	if err := json.Unmarshal([]byte(clean), &raw); err != nil {
		return Response{}
	}

	res := Response{
		NeedExplanation: raw.NeedExplanation == "true",
		Question:        raw.Question,
		NeedPrereqs:     raw.NeedPrereqs == "true",
		Prereqs:         raw.Prereqs,
		Risky:           raw.Risky == "true",
		Commands:        raw.Commands,
		Confidence:      raw.Confidence == "true",
	}

	return res
}
