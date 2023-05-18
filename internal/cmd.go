package internal

import (
	"os/exec"
	"strings"
)

func Command(name string, arg ...string) (string, bool) {
	out, err := exec.Command(name, arg...).Output()
	return strings.TrimSpace(string(out)), err == nil
}
