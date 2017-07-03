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
	"path"

	"github.com/sascha-andres/devenv/helper"
)

type repositoryPullCommand struct{}

func (c repositoryPullCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	_, repository := i.EnvConfiguration.GetRepository(repositoryName)
	if repository.Pinned {
		return fmt.Errorf("Repository %s is pinned. Please unpin if you want to update", repository.Name)
	}
	repositoryPath := path.Join(i.ExecuteScriptDirectory, repository.Path)
	var arguments []string
	arguments = append(arguments, "pull")
	arguments = append(arguments, args...)
	vars, err := i.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	_, err = helper.Git(vars, repositoryPath, arguments...)
	return err
}

func (c repositoryPullCommand) IsResponsible(commandName string) bool {
	return commandName == "pull" || commandName == "<"
}

func init() {
	repositoryCommands = append(repositoryCommands, repositoryPullCommand{})
}
