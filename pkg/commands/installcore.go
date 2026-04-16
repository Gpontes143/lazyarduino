// Package commands
package commands

import "os/exec"

func InstallCore(CoreName string) (string, error) {
	cmd := exec.Command("arduino-cli", "core", "install", CoreName, "--json")
	output, err := cmd.CombinedOutput()
	return string(output), err
}
