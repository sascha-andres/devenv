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

package commit

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_commit"
	"log"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repositoryName string, args []string) error {
	for _, repository := range e.Repositories {
		if repository.Disabled || repository.Pinned != "" {
			continue
		}
		log.Printf("Commit for '%s'\n", repository.Name)
		r := repo_commit.Command{}
		r.Execute(e, executeScriptDirectory, repository.Name, args)
	}
	return nil
}
