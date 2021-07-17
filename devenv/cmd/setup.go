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

// pullCmd represents the pull command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Get a project on your harddrive",
	Long: `Call this to create the directory structure for your multi
repository project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		log.Printf("Called to get '%s'\n", projectName)
		if ok, err := helper2.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
			if "" == projectName || !devenv.ProjectIsCreated(projectName) {
				projectDirectory := path.Join(viper.GetString("basepath"), projectName)
				os.MkdirAll(projectDirectory, 0700)
				var ev devenv.EnvironmentConfiguration
				ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml"))
				for _, repo := range ev.Repositories {
					if repo.Disabled {
						log.Printf("Repository %s is disabled", repo.Name)
						continue
					}
					if _, err := helper2.Git(ev.ProcessConfiguration.Environment, projectDirectory, "clone", repo.URL, repo.Path); err != nil {
						log.Fatalf("Error executing git: '%#v'", err)
					}
					if repo.Pinned != "" {
						if _, err := helper2.Git(ev.ProcessConfiguration.Environment, path.Join(projectDirectory, repo.Path), "checkout", repo.Pinned); err != nil {
							log.Fatalf("Error executing git: '%#v'", err)
						}
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
	RootCmd.AddCommand(setupCmd)
}
