package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ollama struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResp struct {
	Response string `json:"response"`
}

func SendToOllama(prompt string) (string, error) {
	request := Ollama{
		Model:  "llama3.2",
		Prompt: prompt,
		Stream: false,
	}
	marshalled, err := json.Marshal(request)
	if err != nil {
		return " ", fmt.Errorf("Failed In Converting The Request %v", err)
	}
	url := "http://localhost:11434/api/generate"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshalled))
	if err != nil {
		return " ", fmt.Errorf("POST Request failed Check the service %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return " ", fmt.Errorf("Error during reading response %v", err)
	}
	output := OllamaResp{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		return " ", fmt.Errorf("Error during parsing JSON %v", err)
	}

	return output.Response, nil
}
