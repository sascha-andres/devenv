// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
)

// renameCmd represents the clean command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename project environment",
	Long: `Rename an environment. This renames the config file as well as
the directory containing the repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := strings.Join(args, " ")
		parameters := strings.Split(arguments, "->")

		if len(parameters) != 2 {
			log.Print("Usage: devenv oldname -> newname")
			os.Exit(1)
		}

		projectName := strings.TrimSpace(parameters[0])
		newProjectName := strings.TrimSpace(parameters[1])

		if projectName == "" || newProjectName == "" {
			log.Print("You have to provide both project names")
			log.Print("Usage: devenv oldname -> newname")
			os.Exit(1)
		}

		if projectName != "" && devenv.ProjectIsCreated(projectName) {
			configurationFile := path.Join(viper.GetString("configpath"), projectName+".yaml")
			newConfigurationFile := path.Join(viper.GetString("configpath"), newProjectName+".yaml")
			if ok, err := helper.Exists(configurationFile); ok && err == nil {
				if ok, err := helper.Exists(newConfigurationFile); ok && err == nil {
					log.Printf("project with the name %s already exists!", newProjectName)
					os.Exit(1)
				}
				os.Rename(path.Join(viper.GetString("basepath"), projectName), path.Join(viper.GetString("basepath"), newProjectName))
				os.Rename(configurationFile, newConfigurationFile)
			} else {
				log.Printf("project with the name %s does not exist!", newProjectName)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(renameCmd)
}
