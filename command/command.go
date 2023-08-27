package command

import (
	"bytes"
	"os/exec"
	"strings"
)

type Command struct {
	Command string `json:"command"`
}

func (c *Command) RunCommand() (string, error) {
	cmdFields := strings.Fields(c.Command)
	cmd := cmdFields[0]
	args := cmdFields[1:]

	cmdObj := exec.Command(cmd, args...)
	var out bytes.Buffer
	cmdObj.Stdout = &out

	err := cmdObj.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
