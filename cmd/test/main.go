package main

import (
	"fmt"
	"trinethra/internal/ollama"
)

func main() {
	response, err := ollama.SendToOllama("What is 2 + 2?")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	fmt.Println("Response ", response)
}
