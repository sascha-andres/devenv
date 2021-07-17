package os_helper

import (
	"os"
	"os/exec"
)

// GetCommand creates a command structure
func GetCommand(commandName string, env Environ, path string, arguments ...string) (*exec.Cmd, error) {
	command := exec.Command(commandName, arguments...)
	command.Dir = path
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return command, nil
}
