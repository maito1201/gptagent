package main

import (
	"encoding/json"
	"fmt"
	"gptagent/command"
	"gptagent/gpt"
	"os"
)

var g gpt.GPT

func init() {
	url := os.Getenv("OPENAI_API_ENDPOINT")
	if url == "" {
		url = "https://api.openai.com/v1/chat/completions"
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable is not set")
	}
	g = gpt.NewGPT(url, apiKey)
}

func main() {
	arg := os.Args[1]

	messages := gpt.Messages{
		{Role: "system", Content: "You are an agent suggested and instructed to execute a linux command, if you are asked to execute a command please follow the instructions"},
		{Role: "user", Content: "Please tell me the linux command to get the current time"},
		{Role: "assistant", Content: "date"},
		{Role: "user", Content: arg},
	}

	functions := gpt.Functions{
		{
			Name:        "execute_linux_command",
			Description: "input linux command if user need to execute linux command",
			Parameters: gpt.Parameters{
				Type: "object",
				Properties: map[string]gpt.Property{
					"command": {
						Type:        "string",
						Description: `linux command to execute`,
					},
				},
				Required: []string{"command"},
			},
		},
	}

	model := os.Getenv("GPT_MODEL")
	if model == "" {
		model = "gpt-3.5-turbo"
	}

	reqData := gpt.RequestData{
		Model:     model,
		Messages:  messages,
		Functions: functions,
	}

	d, err := g.CallCompletionsAPI(reqData)
	if err != nil {
		panic(err)
	}

	if len(d.Choices) > 0 {
		if d.Choices[0].Message.Content != nil {
			fmt.Println(*d.Choices[0].Message.Content)
		} else if d.Choices[0].Message.FunctionCall != nil {
			if d.Choices[0].Message.FunctionCall.Name == "execute_linux_command" {
				var c command.Command
				if err := json.Unmarshal([]byte(d.Choices[0].Message.Arguments), &c); err != nil {
					panic(err)
				}
				fmt.Printf("running... %s\n", c.Command)
				c.RunCommand()
			} else {
				fmt.Printf("Error: ChatGPT called unknown function %s\n", d.Choices[0].Message.FunctionCall.Name)
			}
		}
	}
}