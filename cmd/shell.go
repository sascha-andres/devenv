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

	"github.com/chzyer/readline"
	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
)

var ev devenv.EnvironmentConfiguration

var completer = readline.NewPrefixCompleter(
	readline.PcItem("repo",
		readline.PcItemDynamic(listRepositories(),
			readline.PcItem("branch"),
			readline.PcItem("commit"),
			readline.PcItem("log"),
			readline.PcItem("pull"),
			readline.PcItem("push"),
			readline.PcItem("status"),
		),
	),
	readline.PcItem("addrepo"),
	readline.PcItem("branch"),
	readline.PcItem("commit"),
	readline.PcItem("delrepo"),
	readline.PcItem("log"),
	readline.PcItem("pull"),
	readline.PcItem("push"),
	readline.PcItem("status"),
	readline.PcItem("quit"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start devenv shell",
	Long:  `Devenv shell allows to work with all repositories at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runInterpreter(args)
		if err != nil {
			log.Fatalf("Error: '%s'", err.Error)
		}
	},
}

func init() {
	RootCmd.AddCommand(shellCmd)
}
