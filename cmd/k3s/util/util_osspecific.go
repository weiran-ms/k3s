//go:build !windows
// +build !windows
package util

import (
	"syscall"
)

const (
	PATH_ENV_SEPARATOR = ":"
)

func Exec(cmd string, args []string) error {
	return syscall.Exec(cmd, args, os.Environ())
}