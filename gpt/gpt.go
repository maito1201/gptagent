package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GPT struct {
	Url    string
	APIKey string
}

func NewGPT(url string, apiKey string) GPT {
	return GPT{
		Url:    url,
		APIKey: apiKey,
	}
}

func (g *GPT) CallCompletionsAPI(rd RequestData) (*ResponseData, error) {
	reqBody, err := json.Marshal(rd)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", g.Url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var d ResponseData
	if err = json.Unmarshal(respBody, &d); err != nil {
		return nil, err
	}

	if len(d.Choices) <= 0 {
		fmt.Println(string(respBody))
	}
	return &d, nil
}
