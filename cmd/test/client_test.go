package test

import (
	"fmt"
	"testing"
	"trinethra/internal/ollama"
)

func TestOllama(t *testing.T) {
	response, err := ollama.SendToOllama("What is 2 + 2?")
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	fmt.Print("Response ", response)
}
