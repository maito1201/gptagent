package command

import (
	"os/exec"
	"strings"
)

type Command struct {
	Command string `json:"command"`
}

func (c *Command) RunCommand() string {
	cmdFields := strings.Fields(c.Command)
	args := cmdFields[0:]
	args = append(args, "2>&1")

	cmd := exec.Command("sh", "-c", strings.Join(args, " "))

	out, _ := cmd.Output()
	return string(out)
}
