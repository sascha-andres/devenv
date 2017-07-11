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

package interactive

import "log"

type commitCommand struct{}

func (c commitCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	for _, repository := range i.EnvConfiguration.Repositories {
		if repository.Disabled || repository.Pinned != "" {
			continue
		}
		log.Printf("Commit for '%s'\n", repository.Name)
		r := repositoryCommitCommand{}
		r.Execute(i, repository.Name, args)
	}
	return nil
}

func (c commitCommand) IsResponsible(commandName string) bool {
	return commandName == "commit" || commandName == "ci"
}

func init() {
	commands = append(commands, commitCommand{})
}