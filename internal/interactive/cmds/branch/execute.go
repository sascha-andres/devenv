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

package branch

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"log"
	"path"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repository string, args []string) error {
	for _, repository := range e.Repositories {
		if repository.Disabled || repository.Pinned != "" {
			continue
		}
		log.Printf("Branch for '%s'\n", repository.Name)
		repositoryPath := path.Join(executeScriptDirectory, repository.Path)
		hasBranch, err := helper.HasBranch(helper.GetProcess(e).Environment, repositoryPath, args[0])
		if err != nil {
			return err
		}
		var arguments []string
		arguments = append(arguments, "checkout")
		if !hasBranch {
			arguments = append(arguments, "-b")
		}
		arguments = append(arguments, args...)
		vars, err := e.GetReplacedEnvironment()
		if err != nil {
			return err
		}
		if _, err = helper.Git(vars, repositoryPath, arguments...); err != nil {
			return err
		}
	}
	return nil
}
