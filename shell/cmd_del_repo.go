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

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
)

type delRepositoryCommand struct{}

func (c delRepositoryCommand) Execute(i *Interpreter, repository string, args []string) error {
	log.Println("Delete a repository")
	log.Print("Please provide name: ")
	name := getAnswer()
	index, repositoryInstance := i.EnvConfiguration.GetRepository(name)
	if nil == repositoryInstance {
		log.Fatalln("Repository not found")
	}
	repositoryPath := path.Join(i.ExecuteScriptDirectory, repositoryInstance.Path)
	changes(i, repositoryPath)
	if err := os.RemoveAll(repositoryPath); err != nil {
		return fmt.Errorf("Error removing repository from disk: '%s'", err.Error())
	}
	i.EnvConfiguration.Repositories = append(i.EnvConfiguration.Repositories[:index], i.EnvConfiguration.Repositories[index+1:]...)
	i.EnvConfiguration.SaveToFile(path.Join(viper.GetString("configpath"), i.EnvConfiguration.Name+".yaml"))
	return nil
}

func changes(i *Interpreter, repoPath string) {
	if ok, err := helper.HasChanges(i.getProcess().Environment, repoPath); ok || err != nil {
		if ok {
			log.Fatalln("Changes found, aborting")
		} else {
			log.Fatalf("Error determining if there are changes: '%s'", err.Error())
		}
	}
}

func (c delRepositoryCommand) IsResponsible(commandName string) bool {
	return commandName == "delrepo"
}

func init() {
	commands = append(commands, delRepositoryCommand{})
}
