package task

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunSingle(input_string, input_script string) (string, error) {
	//TODO: Change this into a config parameter.
	scripts_dir := "."

	if scripts_dir[len(scripts_dir)-1:] != "/" {
		scripts_dir = scripts_dir + "/"
	}

	out, err := exec.Command(scripts_dir + strings.TrimSpace(input_script), strings.TrimSpace(input_string)).Output()
	if err != nil {
		return string(out), fmt.Errorf("failed to execute script: %v", err)
	}

	return string(out), nil
}
