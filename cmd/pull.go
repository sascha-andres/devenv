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
	"log"
	"os"
	"path"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Get a project on your harddrive",
	Long: `Call this to create the directory structure for your multi
repository project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		log.Printf("Called to get '%s'\n", projectName)
		if ok, err := helper.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
			if !devenv.ProjectIsCreated(projectName) {
				projectDirectory := path.Join(viper.GetString("basepath"), projectName)
				os.MkdirAll(projectDirectory, 0700)
				var ev devenv.EnvironmentConfiguration
				ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml"))
				for _, repo := range ev.Repositories {
					if err := helper.Git(ev.Environment, projectDirectory, "clone", repo.URL, repo.Path); err != nil {
						log.Fatalf("Error executing git: '%#v'", err)
					}
				}
			} else {
				log.Println("Project is already pulled")
			}
		} else {
			log.Fatalf("'%s' is not a valid project", projectName)
		}
	},
}

func init() {
	RootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
