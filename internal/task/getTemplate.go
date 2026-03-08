package task

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Template struct {
    Scripts [] map[string]Script `mapstructure:"scripts"`
    Message string `mapstructure:"message"`
}

type Script struct {
	Name string `mapstructure:"script"`
    Args string `mapstructure:"args"`
}

func GetTemplate(passedTemplate *Template, templateName, templateDirectory string) error {
    configDir := filepath.Join(templateDirectory, templateName)
	if !strings.Contains(templateName, ".yaml") {
		configDir += ".yaml"
	}

    viper.SetConfigFile(configDir)
    if err := viper.ReadInConfig(); err != nil {
            return err
    }

    if err := viper.Unmarshal(&passedTemplate); err != nil {
            return err
    }

	return nil
}
