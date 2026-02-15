package task

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Dogel-ai/amici/util"
)

func RunSingle(input_string, input_script string) (string, error) {
	//TODO: Change this into a config parameter.
	scriptDir := "./../mod-scripts"

	input_script = strings.TrimSpace(input_script)

	if scriptDir[len(scriptDir)-1:] != "/" {
		scriptDir += "/"
	}
	scriptDir += input_script
	
	if err := util.IsValidScript(scriptDir); err != nil {
		return "", err
	}

	out, err := exec.Command(scriptDir, input_string).Output()
	if err != nil {
		return string(out), fmt.Errorf("failed to execute script: %v", err)
	}

	return string(out), nil
}
