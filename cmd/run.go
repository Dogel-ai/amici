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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/Dogel-ai/amici/internal/task"
)

var (
	scriptChoice []string
	templateChoice string
	sumOutput string
)

var runCmd = &cobra.Command{
	Use:   "run [string]",
	Short: "Converts input through a modular selection of script",
	Long:  `Converts input through a modular selection of scripts.

This tool takes an argument as input and passes it through a selection of
scripts within templates, or through singular scripts selected with the --script/-s flag.

Examples:
  amici run Example String			using the default template
  amici run Example String -s exampleScript.py	using only exampleScript.py`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if(scriptChoice != nil) {
			for _, v := range scriptChoice {
				out, err := task.RunSingle(args[0], v)
				if err != nil {
					return err
				}
				sumOutput += out
			}
			fmt.Println(sumOutput)
		} else {
			var template task.Template
			err := task.GetTemplate(&template, templateChoice, viper.GetString("templates_directory"))
			if err != nil {
				return err
			}
			
			finalMessage := template.Message

			for _, script := range template.Scripts {
				for key, i := range script {
					cmd := strings.Replace(i.Args, "%i", args[0], -1)
					out, err := task.RunSingle(cmd, i.Name)
					if err != nil {
						return err
					}
					finalMessage = strings.Replace(finalMessage, "%" + key, out, -1)
				}
			}
			fmt.Println(finalMessage)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringSliceVarP(&scriptChoice, "script", "s", nil, "select singular script(s) to run")
	runCmd.Flags().StringVarP(&templateChoice, "template", "t", "defaultTemplate", "select a template to use")
}
