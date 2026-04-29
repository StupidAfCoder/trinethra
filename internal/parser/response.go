package parser

import (
	"encoding/json"
	"fmt"
	"strings"
	"trinethra/package/models"
)

func ParseRespone(raw_resp string) (models.Analysis, error) {
	firstIndex := strings.Index(raw_resp, "{")
	lastIndex := strings.LastIndex(raw_resp, "}")
	if firstIndex == -1 || lastIndex == -1 {
		return models.Analysis{}, fmt.Errorf("LLM response was not up to mark Empty Analysis")
	}
	json_str := raw_resp[firstIndex : lastIndex+1]
	var analysis models.Analysis
	err := json.Unmarshal([]byte(json_str), &analysis)
	if err != nil {
		return models.Analysis{}, fmt.Errorf("Error during Parsing Json %v", err)
	}
	return analysis, nil
}
