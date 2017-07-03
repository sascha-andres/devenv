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

type repositoryCommitCommand struct{}

func (c repositoryCommitCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	_, repository := i.EnvConfiguration.GetRepository(repositoryName)
	if repository.Disabled || repository.Pinned != "" {
		return nil
	}
	repositoryPath := path.Join(i.ExecuteScriptDirectory, repository.Path)
	if hasChanges, err := helper.HasChanges(i.getProcess().Environment, repositoryPath); hasChanges && err == nil {
		if _, err := helper.Git(i.getProcess().Environment, repositoryPath, "add", "--all", ":/"); err != nil {
			return err
		}
		if err := execHelper(i, repositoryPath, "commit", args); err != nil {
			return err
		}
	} else {
		fmt.Println(" --> no changes")
	}
	return nil
}

func (c repositoryCommitCommand) IsResponsible(commandName string) bool {
	return commandName == "commit" || commandName == "ci"
}

func init() {
	repositoryCommands = append(repositoryCommands, repositoryCommitCommand{})
}
