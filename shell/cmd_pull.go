// Licensed under the Apache License, Version 2.0 (the "License");
// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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
	"log"
	"path"

	"github.com/sascha-andres/devenv/helper"
)

type pullCommand struct{}

func (c pullCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	for _, repository := range i.EnvConfiguration.Repositories {
		if repository.Disabled || repository.Pinned {
			continue
		}
		log.Printf("Pull for '%s'\n", repository.Name)
		repositoryPath := path.Join(i.ExecuteScriptDirectory, repository.Path)
		var arguments []string
		arguments = append(arguments, "pull")
		arguments = append(arguments, args...)
		vars, err := i.EnvConfiguration.GetReplacedEnvironment()
		if err != nil {
			return err
		}
		if _, err = helper.Git(vars, repositoryPath, arguments...); err != nil {
			return err
		}
	}
	return nil
}

func (c pullCommand) IsResponsible(commandName string) bool {
	return commandName == "pull" || commandName == "<"
}

func init() {
	commands = append(commands, pullCommand{})
}
