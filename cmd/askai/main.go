package main

import (
	"fmt"
	"os"
	"strings"

	"ask-ai/internal/ollama"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: askai \"your question\"")
		return
	}

	prompt := strings.Join(os.Args[1:], " ")

	answer, err := ollama.Generate(prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(answer)
}
