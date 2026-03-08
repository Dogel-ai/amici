package task

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Dogel-ai/amici/util"
	"github.com/spf13/viper"
)

func RunSingle(inputString, inputScript string) (string, error) {
	scriptDir := viper.GetString("scripts_directory")

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

	if out[len(out)-1] == 10 {
		return string(out[0:len(out)-1]), nil
	}

	return string(out), nil
}
