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
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/chzyer/readline"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/sascha-andres/devenv/shell"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("repo"),
	readline.PcItem("addrepo"),
	readline.PcItem("branch"),
	readline.PcItem("commit"),
	readline.PcItem("delrepo"),
	readline.PcItem("log"),
	readline.PcItem("pull"),
	readline.PcItem("push"),
	readline.PcItem("status"),
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start devenv shell",
	Long:  `Devenv shell allows to work with all repositories at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		log.Printf("Called to start shell for '%s'\n", projectName)
		if "" == projectName || !devenv.ProjectIsCreated(projectName) {
			log.Fatalf("Project '%s' does not yet exist", projectName)
		}
		var ev devenv.EnvironmentConfiguration
		if ok, err := helper.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
			if err := ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml")); err != nil {
				log.Fatalf("Error reading env config: '%s'", err.Error())
			}
		}

		reader := bufio.NewReader(os.Stdin)
		interp := shell.NewInterpreter(path.Join(viper.GetString("basepath"), projectName), ev)
		for {
			fmt.Print("> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error getting command: %#v", err)
			}
			text = strings.TrimSpace(text)
			if "quit" == text || "q" == text {
				break
			}
			if err := interp.Execute(text); err != nil {
				log.Printf("Error: '%s'", err.Error())
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(shellCmd)
}
