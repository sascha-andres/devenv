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

package del_repo

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repository string, args []string) error {
	log.Println("Delete a repository")
	log.Print("Please provide name: ")
	name := helper.GetAnswer()
	index, repositoryInstance := e.GetRepository(name)
	if nil == repositoryInstance {
		log.Fatalln("Repository not found")
	}
	repositoryPath := path.Join(executeScriptDirectory, repositoryInstance.Path)
	changes(e, repositoryPath)
	if err := os.RemoveAll(repositoryPath); err != nil {
		return errors.Wrap(err, "could not remove repository from disk")
	}
	e.Repositories = append(e.Repositories[:index], e.Repositories[index+1:]...)
	e.SaveToFile(path.Join(viper.GetString("configpath"), e.Name+".yaml"))
	return nil
}
