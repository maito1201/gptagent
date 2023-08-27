package command

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Command struct {
	Command string `json:"command"`
}

func (c *Command) RunCommand() {
	cmdFields := strings.Fields(c.Command)
	cmd := cmdFields[0]
	args := cmdFields[1:]

	cmdObj := exec.Command(cmd, args...)
	var out bytes.Buffer
	cmdObj.Stdout = &out

	err := cmdObj.Run()
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	fmt.Printf("Command output: \n%s\n", out.String())
}
