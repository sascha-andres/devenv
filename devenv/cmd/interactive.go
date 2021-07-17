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
	"fmt"
	helper2 "github.com/sascha-andres/devenv/internal/helper"
	interactive2 "github.com/sascha-andres/devenv/internal/interactive"
	"io"
	"log"
	"path"
	"strings"

	"os"

	"github.com/chzyer/readline"
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/spf13/viper"
)

var ev devenv.EnvironmentConfiguration

var completer = readline.NewPrefixCompleter(
	readline.PcItem("repo",
		readline.PcItemDynamic(listRepositories(),
			readline.PcItem("branch"),
			readline.PcItem("commit"),
			readline.PcItem("log"),
			readline.PcItem("shell"),
			readline.PcItem("pull"),
			readline.PcItem("push"),
			readline.PcItem("status"),
			readline.PcItem("pin"),
			readline.PcItem("unpin"),
			readline.PcItem("enable"),
			readline.PcItem("disable"),
		),
	),
	readline.PcItem("addrepo"),
	readline.PcItem("addscript"),
	readline.PcItem("editscript"),
	readline.PcItem("delscript"),
	readline.PcItem("branch"),
	readline.PcItem("commit"),
	readline.PcItem("delrepo"),
	readline.PcItem("log"),
	readline.PcItem("pull"),
	readline.PcItem("push"),
	readline.PcItem("status"),
	readline.PcItem("quit"),
	readline.PcItem("scan"),
	readline.PcItem("shell"),
)

//func filterInput(r rune) (rune, bool) {
//	switch r {
//	// block CtrlZ feature
//	case readline.CharCtrlZ:
//		return r, false
//	}
//	return r, true
//}

func setup(projectName string) error {
	if "" == projectName || !devenv.ProjectIsCreated(projectName) {
		return errors.New(fmt.Sprintf("Project '%s' does not yet exist", projectName))
	}
	if ok, err := helper2.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
		if err := ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml")); err != nil {
			return errors.New(fmt.Sprintf("Error reading env config: '%s'", err.Error()))
		}
	}
	return nil
}

func runInterpreter(args []string) error {
	projectName := strings.Join(args, " ")
	log.Printf("Called to start shell for '%s'", projectName)
	if "" == projectName {
		os.Exit(1)
	}
	err := setup(projectName)
	if err != nil {
		return err
	}

	interpreter := interactive2.NewInterpreter(path.Join(viper.GetString("basepath"), projectName), ev)
	l, err := getReadlineConfig(projectName)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			log.Fatalf("Error closing readline: " + err.Error())
		}
	}()

	log.SetOutput(l.Stderr())

	for {
		line, doBreak := getLine(l)
		if doBreak {
			break
		}
		line = strings.TrimSpace(line)
		switch line {
		case "quit", "q":
			return nil
		default:
			executeLine(interpreter, line)
			if err := setup(projectName); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func getLine(l *readline.Instance) (string, bool) {
	line, err := l.Readline()
	if err == readline.ErrInterrupt {
		if len(line) == 0 {
			return "", true
		}
		return line, false
	} else if err == io.EOF {
		return "", true
	}
	return line, false
}

func executeLine(interpreter *interactive2.Interpreter, line string) {
	err := interpreter.Execute(line)
	if err != nil {
		log.Println(err)
	}
}
func getReadlineConfig(projectName string) (*readline.Instance, error) {
	return readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/devenv-" + projectName + ".tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
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
