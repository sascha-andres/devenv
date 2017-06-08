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

type commitCommand struct{}

func (c commitCommand) Execute(i *Interpreter, repository string, args []string) error {
	for _, repo := range i.EnvConfiguration.Repositories {
		if repo.Disabled {
			continue
		}
		fmt.Printf("Commit for '%s'\n", repo.Name)
		repoPath := path.Join(i.ExecuteScriptDirectory, repo.Path)
		if hasChanges, err := helper.HasChanges(i.EnvConfiguration.Environment, repoPath); hasChanges && err == nil {
			if _, err := helper.Git(i.EnvConfiguration.Environment, repoPath, "add", "--all", ":/"); err != nil {
				return err
			}
			var arguments []string
			arguments = append(arguments, "commit")
			arguments = append(arguments, args...)
			_, err := helper.Git(i.EnvConfiguration.Environment, repoPath, arguments...)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(" --> no changes")
		}
	}
	return nil
}

func (c commitCommand) IsResponsible(commandName string) bool {
	return commandName == "commit" || commandName == "ci"
}

func init() {
	commands = append(commands, commitCommand{})
}
