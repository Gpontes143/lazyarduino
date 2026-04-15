package commands

import "os/exec"

func Upload(port string, fqbn string, path string) (string, error) {
	cmd := exec.Command("arduino-cli", "Upload", "-p", port, "--fqbn", fqbn, path)

	output, err := cmd.CombinedOutput()
	return string(output), err
}
