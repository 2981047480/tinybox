package exec_command

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Command interface {
	Exec() (string, error)
}

type LocalCommand struct {
	Command string
}

func NewLocalCommand(command string) LocalCommand {
	return LocalCommand{
		Command: command,
	}
}

func (c *LocalCommand) Exec() (string, error) {
	cmd := exec.Command("bash", "-c", c.Command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command exec failed, err: %v", err)
	}
	return out.String(), nil
}
