package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type Response struct {
	Response string `json:"response"` // matches Ollama API
}

// Generate sends a prompt to Ollama and returns the response
func Generate(prompt string) (string, error) {
	req := Request{
		Model:  "kimi-k2.5:cloud", // safe small model
		Prompt: prompt,
		Stream: false,
	}

	body, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Ollama API error: %s", resp.Status)
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}