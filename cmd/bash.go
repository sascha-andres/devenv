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
	"path"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// bashCmd represents the bash command
var (
	bashCmd = &cobra.Command{
		Use:   "bash",
		Short: "Start a bash shell",
		Long:  `Spwans a bash shell in the directory where all code is located`,
		Run: func(cmd *cobra.Command, args []string) {
			projectName := strings.Join(args, " ")
			log.Printf("Called to start shell for '%s'\n", projectName)
			if !devenv.ProjectIsCreated(projectName) {
				log.Fatalf("Project '%s' does not yet exist", projectName)
			}
			projectFileNamePath := path.Join(viper.GetString("configpath"), projectName+".yaml")
			log.Printf("Loading from '%s'\n", projectFileNamePath)

			var ev devenv.EnvironmentConfiguration
			if err := ev.LoadFromFile(projectFileNamePath); err != nil {
				log.Fatalf("Error loading environment config: %#v\n", err)
			}
			if err := ev.StartShell(); err != nil {
				log.Fatalf("Error starting shell: %#v\n", err)
			}
		},
	}
)

func init() {
	RootCmd.AddCommand(bashCmd)
}
