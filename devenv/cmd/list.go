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
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
)

// listCmd represents the list command which prints information
// about all repositories in a development environment
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all repositories within an environment",
	Long:  `List`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		if "" == projectName || !devenv.ProjectIsCreated(projectName) {
			log.Fatalf("Project '%s' does not exist", projectName)
			os.Exit(1)
		}
		projectFileNamePath := path.Join(viper.GetString("configpath"), projectName+".yaml")
		if ok, err := helper.Exists(projectFileNamePath); !ok || err != nil {
			log.Fatalf("'%s' does not exist", projectFileNamePath)
			os.Exit(1)
		}

		var ev devenv.EnvironmentConfiguration
		if err := ev.LoadFromFile(projectFileNamePath); err != nil {
			log.Fatalf("Error loading environment config: %#v\n", err)
			os.Exit(1)
		}
		for i := range ev.Repositories {
			log.Printf("%s: ", ev.Repositories[i].Name)
			log.Printf("  path: %s: ", ev.Repositories[i].Path)
			log.Printf("  remote: %s", ev.Repositories[i].URL)
			log.Printf("  is disabled: %v", ev.Repositories[i].Disabled)
			log.Printf("  is pinned to: %s", ev.Repositories[i].Pinned)
		}
	},
}

func init() {
	// flags
	RootCmd.AddCommand(listCmd)
}
