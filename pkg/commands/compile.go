package commands

import "os/exec"

func Compile(fqbn string, path string) (string, error) {
	cmd := exec.Command("arduino-cli", "compile", "--fqbn", fqbn, path)

	output, err := cmd.CombinedOutput()
	return string(output), err
}
