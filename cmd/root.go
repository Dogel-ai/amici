/*
Copyright © 2026 Dogel <dogel.kszb@proton.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configDir string

var rootCmd = &cobra.Command{
	Use:   "amici",
	Short: "Converts input through a modular selection of scripts.",
	Long: `Converts input through a modular selection of scripts.

This tool takes an argument as input and passes it through a selection of
scripts within templates, or through singular scripts selected with the --script/-s flag.

Examples:
  amici run Example String                              using the default template
  amici run Example String -s exampleScript.py          using only exampleScript.py`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configDir, "config", "", "config file (default is $HOME/.config/amici)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if configDir != "" {
		viper.SetConfigFile(configDir)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to find home directory")
		}
		configDir = home + "/.config/amici"

		viper.AddConfigPath(configDir)
		viper.SetConfigName("amici")
		viper.SetConfigType("yaml")
	}
	
	if err := viper.ReadInConfig(); err != nil {
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			fmt.Fprintln(os.Stderr, "Could not create config directory ", configDir)
			os.Exit(1)
		}
		if err := viper.SafeWriteConfig(); err != nil {
			fmt.Fprintln(os.Stderr, "Could not create config file within ", configDir)
			os.Exit(1)
		}
	}
}
