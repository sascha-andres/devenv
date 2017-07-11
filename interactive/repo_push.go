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

package interactive

import (
	"path"

	"github.com/sascha-andres/devenv/helper"
)

type repoPushCommand struct{}

func (c repoPushCommand) Execute(i *Interpreter, repository string, args []string) error {
	_, repo := i.EnvConfiguration.GetRepository(repository)
	repoPath := path.Join(i.ExecuteScriptDirectory, repo.Path)
	var arguments []string
	arguments = append(arguments, "push")
	arguments = append(arguments, args...)
	vars, err := i.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	_, err = helper.Git(vars, repoPath, arguments...)
	return err
}

func (c repoPushCommand) IsResponsible(commandName string) bool {
	return commandName == "push" || commandName == ">"
}

func init() {
	repositoryCommands = append(repositoryCommands, repoPushCommand{})
}