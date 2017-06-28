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

type repoCommitCommand struct{}

func (c repoCommitCommand) Execute(i *Interpreter, repository string, args []string) error {
	_, repo := i.EnvConfiguration.GetRepository(repository)
	repoPath := path.Join(i.ExecuteScriptDirectory, repo.Path)
	if hasChanges, err := helper.HasChanges(i.EnvConfiguration.Environment, repoPath); hasChanges && err == nil {
		if _, err := helper.Git(i.EnvConfiguration.Environment, repoPath, "add", "--all", ":/"); err != nil {
			return err
		}
		if err := execHelper(i, repoPath, "commit", args); err != nil {
			return err
		}
	} else {
		fmt.Println(" --> no changes")
	}
	return nil
}

func (c repoCommitCommand) IsResponsible(commandName string) bool {
	return commandName == "commit" || commandName == "ci"
}

func init() {
	repositoryCommands = append(repositoryCommands, repoCommitCommand{})
}
