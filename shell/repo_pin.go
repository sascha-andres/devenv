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
	"log"
)

type repositoryPinCommand struct{}

func (c repositoryPinCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	_, repository := i.EnvConfiguration.GetRepository(repositoryName)
	log.Printf("Pinning %s", repository.Name)
  // git rev-parse --verify HEAD
	return nil
}

func (c repositoryPinCommand) IsResponsible(commandName string) bool {
	return commandName == "pin"
}

func init() {
	repositoryCommands = append(repositoryCommands, repositoryPinCommand{})
}
