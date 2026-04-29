package handler

import (
	"encoding/json"
	"net/http"
	"trinethra/internal/ollama"
	"trinethra/internal/prompt"
	"trinethra/internal/response"
)

type AnalyzeRequest struct {
	Transcript string `json:"transcript"`
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var analyze AnalyzeRequest
	err := json.NewDecoder(r.Body).Decode(&analyze)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if analyze.Transcript == "" {
		http.Error(w, "Empty Transcript Was Passed!", http.StatusBadRequest)
		return
	}

	prompt_str := prompt.ReturnPrompt(analyze.Transcript)

	raw_resp, err := ollama.SendToOllama(prompt_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	analysis, err := response.ParseRespone(raw_resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analysis)
}
