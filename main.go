package main

import (
	"bufio"
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

var prompt = `
You are an agent suggested and instructed to execute a linux command,
If you are asked to execute a command please follow the instructions.
If you reply with a Linux command, the user will return the execution result.
Then when you decide there is a command to run next, run the following command.
If you are asked to read something use cat command.
If you are asked to write something use printf command with pipe to output file.
Do not use echo command to write file avoid to unexpected result.
`

func main() {
	arg := os.Args[1]
	messages := gpt.Messages{
		{Role: "system", Content: prompt},
		{Role: "user", Content: "Please tell me the linux command to get the current time"},
		{Role: "assistant", Content: "date"},
		{Role: "user", Content: "Please write shell script to get current time here"},
		{Role: "assistant", Content: "printf date > test.sh"},
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
	execute(&reqData)
}

func execute(reqData *gpt.RequestData) {
	d, err := g.CallCompletionsAPI(reqData)
	if err != nil {
		panic(err)
	}

	if len(d.Choices) > 0 {
		if d.Choices[0].Message.Content != nil {
			reader := bufio.NewReader(os.Stdin)
			ans := *d.Choices[0].Message.Content
		TOP:
			fmt.Println(ans)
			fmt.Print("(input your answer): ")
			input, _ := reader.ReadString('\n')
			if input == "\n" {
				goto TOP
			} else {
				reqData.Messages = append(
					reqData.Messages,
					gpt.Messages{
						{Role: "assistant", Content: ans},
						{Role: "user", Content: input},
					}...,
				)
				execute(reqData)
			}
		} else if d.Choices[0].Message.FunctionCall != nil {
			if d.Choices[0].Message.FunctionCall.Name == "execute_linux_command" {
				var c command.Command
				if err := json.Unmarshal([]byte(d.Choices[0].Message.Arguments), &c); err != nil {
					panic(err)
				}
				out := c.RunCommand()
				reqData.Messages = append(reqData.Messages, gpt.Messages{
					{Role: "assistant", Content: c.Command},
					{Role: "user", Content: out},
				}...)
				execute(reqData)
			} else {
				fmt.Printf("Error: ChatGPT called unknown function %s\n", d.Choices[0].Message.FunctionCall.Name)
			}
		}
	}
}
