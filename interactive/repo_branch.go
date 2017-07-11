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

type repoBranchCommand struct{}

func (c repoBranchCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	_, repository := i.EnvConfiguration.GetRepository(repositoryName)
	if repository.Disabled || repository.Pinned != "" {
		return nil
	}
	repositoryPath := path.Join(i.ExecuteScriptDirectory, repository.Path)
	hasBranch, err := helper.HasBranch(i.getProcess().Environment, repositoryPath, args[0])
	if err != nil {
		return err
	}
	var arguments []string
	arguments = append(arguments, "checkout")
	if !hasBranch {
		arguments = append(arguments, "-b")
	}
	arguments = append(arguments, args...)
	vars, err := i.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	if _, err = helper.Git(vars, repositoryPath, arguments...); err != nil {
		return err
	}
	return nil
}

func (c repoBranchCommand) IsResponsible(commandName string) bool {
	return commandName == "branch" || commandName == "br"
}

func init() {
	repositoryCommands = append(repositoryCommands, repoBranchCommand{})
}