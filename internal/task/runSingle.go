package task

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Dogel-ai/amici/util"
)

func RunSingle(inputString, inputScript string) (string, error) {
	//TODO: Change this into a config parameter.
	scriptDir := "./../mod-scripts"

	inputScript = strings.TrimSpace(inputScript)

	if scriptDir[len(scriptDir)-1:] != "/" {
		scriptDir += "/"
	}
	scriptDir += inputScript
	
	if err := util.IsValidScript(scriptDir); err != nil {
		return "", err
	}

	out, err := exec.Command(scriptDir, inputString).Output()
	if err != nil {
		return string(out), fmt.Errorf("failed to execute script: %v", err)
	}

	return string(out), nil
}
