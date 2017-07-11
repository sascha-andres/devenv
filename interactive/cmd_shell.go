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

type shellCommand struct{}

func (c shellCommand) Execute(i *Interpreter, repository string, args []string) error {
	return i.EnvConfiguration.StartShell()
}

func (c shellCommand) IsResponsible(commandName string) bool {
	return commandName == "shell"
}

func init() {
	commands = append(commands, shellCommand{})
}