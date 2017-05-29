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
	"fmt"
	"path"

	"github.com/sascha-andres/devenv/helper"
)

type pullCommand struct{}

func (c pullCommand) Execute(i *Interpreter, repository string, args []string) error {
	for _, repo := range i.EnvConfiguration.Repositories {
		fmt.Printf("Pull for '%s'\n", repo.Name)
		repoPath := path.Join(i.ExecuteScriptDirectory, repo.Path)
		var arguments []string
		arguments = append(arguments, "pull")
		arguments = append(arguments, args...)
		if err := helper.Git(i.EnvConfiguration.Environment, repoPath, arguments...); err != nil {
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
