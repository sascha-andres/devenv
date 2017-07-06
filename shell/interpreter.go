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
	"strings"

	"github.com/mgutz/str"
	"github.com/pkg/errors"
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
	repositoryCommands []Commander
	commands           []Commander
)

// NewInterpreter returns a new interpreter
func NewInterpreter(path string, ev devenv.EnvironmentConfiguration) *Interpreter {
	return &Interpreter{ExecuteScriptDirectory: path, EnvConfiguration: ev}
}

func (i *Interpreter) getProcess() devenv.EnvironmentExternalProcessConfiguration {
	return i.EnvConfiguration.ProcessConfiguration
}

// Execute takes a line entered by the user and calls the command
func (i *Interpreter) Execute(commandline string) error {
	if strings.TrimSpace(commandline) == "" {
		return nil
	}
	tokenize := str.ToArgv(commandline)
	switch tokenize[0] {
	case "repo":
		return i.executeFromCommands(repositoryCommands, true, tokenize[1:])
	}
	return i.executeFromCommands(commands, false, tokenize)
}

func (i *Interpreter) executeFromCommands(commands []Commander, specific bool, arguments []string) error {
	var err error
	for _, val := range commands {
		var responsible bool
		responsible, err = tryExecuteCommand(val, i, specific, arguments)
		if responsible {
			return err
		}
	}
	if err != nil {
		return err
	}
	if specific {
		return errors.New("'" + arguments[1] + "' is not a valid function")
	}
	return errors.New("'" + arguments[0] + "' is not a valid function")
}

func tryExecuteCommand(val Commander, i *Interpreter, specific bool, arguments []string) (bool, error) {
	commandIndex := 0
	if specific {
		commandIndex = 1
	}
	if val.IsResponsible(arguments[commandIndex]) {
		if specific {
			return true, val.Execute(i, arguments[0], arguments[2:])
		}
		return true, val.Execute(i, "%", arguments[1:])
	}
	return false, nil
}
