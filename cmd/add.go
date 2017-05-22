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
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create e new project environment",
	Long: `Creates a new environment, adding a YAML file in
your environment_config_path.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		log.Printf("Called to add '%s'\n", projectName)
		projectFileNamePath := path.Join(viper.GetString("configpath"), projectName+".yaml")
		log.Printf("Storing in '%s'\n", projectFileNamePath)

		ev := devenv.EnvironmentConfiguration{BranchPrefix: "", Name: projectName}
		result, err := yaml.Marshal(ev)
		if err != nil {
			log.Fatalf("Error marshalling new config: %#v", err)
		}
		if err = ioutil.WriteFile(projectFileNamePath, result, 0600); err != nil {
			log.Fatalf("Error writing new config: %#v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
