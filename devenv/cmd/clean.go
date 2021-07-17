// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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
	helper2 "github.com/sascha-andres/devenv/internal/helper"
	"log"
	"os"
	"path"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean project environment",
	Long: `Removes all code from your harddisk. Scans all repositories for uncommitted
changes before deleting the complete directory tree.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		if projectName != "" && devenv.ProjectIsCreated(projectName) {
			if viper.GetBool("clean.force") {
				log.Print("Forced removal, all changes are lost")
				os.RemoveAll(path.Join(viper.GetString("basepath"), projectName))
				return
			}
			var ev devenv.EnvironmentConfiguration
			if ok, err := helper2.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
				if err := ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml")); err != nil {
					log.Fatalf("Error reading env config: '%s'", err.Error())
				}
			}
			for _, repo := range ev.Repositories {
				if hasChanges, err := helper2.HasChanges(ev.ProcessConfiguration.Environment, path.Join(path.Join(viper.GetString("basepath"), projectName, repo.Path))); hasChanges || err != nil {
					log.Printf("'%s' has changes", repo.Name)
					os.Exit(1)
				}
			}
			log.Printf("'%s' has no changes, removing", projectName)
			os.RemoveAll(path.Join(viper.GetString("basepath"), projectName))
		} else {
			log.Println("Project does not exist")
		}
	},
}

func init() {
	cleanCmd.PersistentFlags().BoolP("force", "f", false, "Force removal")
	viper.BindPFlag("clean.force", cleanCmd.PersistentFlags().Lookup("force"))
	RootCmd.AddCommand(cleanCmd)
}
