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
	"io"
	"log"
	"path"
	"strings"

	"github.com/chzyer/readline"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/sascha-andres/devenv/shell"
	"github.com/spf13/viper"
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

func runInterpreter(args []string) error {
	projectName := strings.Join(args, " ")
	log.Printf("Called to start shell for '%s'\n", projectName)
	if "" == projectName || !devenv.ProjectIsCreated(projectName) {
		log.Fatalf("Project '%s' does not yet exist", projectName)
	}
	if ok, err := helper.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
		if err := ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml")); err != nil {
			log.Fatalf("Error reading env config: '%s'", err.Error())
		}
	}

	interp := shell.NewInterpreter(path.Join(viper.GetString("basepath"), projectName), ev)
	l, err := getReadlineConfig(projectName)
	if err != nil {
		return err
	}
	defer l.Close()

	log.SetOutput(l.Stderr())

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch line {
		case "quit", "q":
			return nil
		default:
			err := interp.Execute(line)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func getReadlineConfig(projectName string) (*readline.Instance, error) {
	return readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/devenv-" + projectName + ".tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
}

func listRepositories() func(string) []string {
	return func(line string) []string {
		var repositories []string
		for _, val := range ev.Repositories {
			if !val.Disabled {
				repositories = append(repositories, val.Name)
			}
		}
		return repositories
	}
}
