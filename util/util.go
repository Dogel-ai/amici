package util

import (
	"os"
	"fmt"
)

func IsValidScript(scriptPath string) error {
        scriptFile, err := os.Stat(scriptPath)
        if err != nil {
                if os.IsNotExist(err) {
                        return fmt.Errorf("file %q not found", scriptPath)
                }
                return fmt.Errorf("failed to read file %q: %v", scriptPath, err)
        }
        if scriptFile.IsDir() {
                return fmt.Errorf("path %q is a directory", scriptPath)
        }
        return nil
}
