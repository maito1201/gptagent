package command

import (
	"os/exec"
)

type Command struct {
	Command string `json:"command"`
}

func (c *Command) RunCommand() string {
	cmd := exec.Command("sh", "-c", c.Command, "2>&1")

	out, _ := cmd.Output()
	return string(out)
}
