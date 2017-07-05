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

import "log"

type logCommand struct{}

func (c logCommand) Execute(i *Interpreter, repository string, args []string) error {
	for _, repo := range i.EnvConfiguration.Repositories {
		if repo.Disabled {
			continue
		}
		log.Printf("Log for '%s'\n", repo.Name)
		var params = []string{"-n", "10"}
		params = append(params, args...)
		var r repositoryLogCommand
		r.Execute(i, repo.Name, params)
	}
	return nil
}

func (c logCommand) IsResponsible(commandName string) bool {
	return commandName == "log" || commandName == "l"
}

func init() {
	commands = append(commands, logCommand{})
}
