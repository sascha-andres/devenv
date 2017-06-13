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

package shell

import (
	"fmt"

	"github.com/mgutz/str"
	"github.com/sascha-andres/devenv"
)

type (
	// Interpreter contains data for where and what to do
	Interpreter struct {
		ExecuteScriptDirectory string
		EnvConfiguration       devenv.EnvironmentConfiguration
	}
)

var (
	repoCommands []Commander
	commands     []Commander
)

// NewInterpreter returns a new interpreter
func NewInterpreter(path string, ev devenv.EnvironmentConfiguration) *Interpreter {
	return &Interpreter{ExecuteScriptDirectory: path, EnvConfiguration: ev}
}

// Execute takes a line entered by the user and calls the command
func (i *Interpreter) Execute(commandline string) error {
	tokenized := str.ToArgv(commandline)
	switch tokenized[0] {
	case "repo", "r":
		return i.executeFromCommands(repoCommands, true, tokenized[1:])
	}
	return i.executeFromCommands(commands, false, tokenized)
}

func (i *Interpreter) executeFromCommands(commands []Commander, specific bool, arguments []string) error {
	commandIndex := 0
	if specific {
		commandIndex = 1
	}
	for _, val := range commands {
		if val.IsResponsible(arguments[commandIndex]) {
			if specific {
				return val.Execute(i, arguments[0], arguments[2:])
			}
			return val.Execute(i, "%", arguments[1:])
		}
	}
	return fmt.Errorf("'%s' is not a valid function", arguments[0])
}
