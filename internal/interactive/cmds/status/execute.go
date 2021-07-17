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

package status

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/interactive/cmds/repo_status"
	"log"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repository string, args []string) error {
	for _, repo := range e.Repositories {
		if repo.Disabled {
			continue
		}
		log.Printf("Status for '%s'\n", repo.Name)
		r := repo_status.Command{}
		err := r.Execute(e, executeScriptDirectory, repo.Name, args)
		if err != nil {
			log.Printf("error printing status of repository: %v", err)
		}
	}
	return nil
}
