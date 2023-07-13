//go:build windows
// +build windows

package util

import (
	"os"
	"os/exec"
)

const (
	PATH_ENV_SEPARATOR = ";"
)

func Exec(cmd string, args []string) error {
	command := exec.Command(cmd, args[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Start()
}
