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
	"path"

	"github.com/spf13/viper"
)

type repositoryUnpinCommand struct{}

func (c repositoryUnpinCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	index, repository := i.EnvConfiguration.GetRepository(repositoryName)
	if repository.Pinned == "" {
		log.Printf("Already unpinned")
		return nil
	}
	log.Printf("Unpinning %s", repository.Name)
	repository.Pinned = ""
	i.EnvConfiguration.Repositories = append(i.EnvConfiguration.Repositories[:index], i.EnvConfiguration.Repositories[index+1:]...)
	i.EnvConfiguration.Repositories = append(i.EnvConfiguration.Repositories, *repository)
	return i.EnvConfiguration.SaveToFile(path.Join(viper.GetString("configpath"), i.EnvConfiguration.Name+".yaml"))
}

func (c repositoryUnpinCommand) IsResponsible(commandName string) bool {
	return commandName == "unpin"
}

func init() {
	repositoryCommands = append(repositoryCommands, repositoryUnpinCommand{})
}
