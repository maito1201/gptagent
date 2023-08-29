package command

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	Command string `json:"command"`
}

func (c *Command) RunCommand() string {

	reader := bufio.NewReader(os.Stdin)
ASK:
	fmt.Printf(`suggested commad is "%s". are you sure to execute? (y/n/input your opinion): `, c.Command)
	input, _ := reader.ReadString('\n')
	if input == "y\n" {
		cmd := exec.Command("sh", "-c", c.Command, "2>&1")
		out, _ := cmd.Output()
		return string(out)
	} else if input == "n\n" {
		return "User does not want to use this command."
	} else if input == "\n" {
		goto ASK
	} else {
		return input
	}
}
