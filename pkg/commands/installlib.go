package commands

import "os/exec"

func InstallLib(Libname string) (string, error) {
	cmd := exec.Command("arduino-cli", "lib", "install", Libname, "--json")

	output, err := cmd.CombinedOutput()
	return string(output), err
}
