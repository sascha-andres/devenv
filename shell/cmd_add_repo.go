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
	"log"
	"path"

	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
)

type addRepositoryCommand struct{}

func (c addRepositoryCommand) Execute(i *Interpreter, repository string, args []string) error {
	log.Println("Add a repo")
	log.Print("Please provide a link: ")
	link := getAnswer()
	log.Print("Please provide a name: ")
	name := getAnswer()
	log.Print("Please provide a relative path: ")
	rpath := getAnswer()
	repositoryPath := path.Join(i.ExecuteScriptDirectory, rpath)
	vars, err := i.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	if _, err := helper.Git(vars, i.ExecuteScriptDirectory, "clone", link, repositoryPath); err != nil {
		log.Fatalf("Error executing git: '%#v'", err)
	}
	rc := devenv.RepositoryConfiguration{
		Path:     rpath,
		Pinned:   "",
		Disabled: false,
		Name:     name,
		URL:      link,
	}
	log.Println(rc)
	i.EnvConfiguration.Repositories = append(i.EnvConfiguration.Repositories, rc)
	return i.EnvConfiguration.SaveToFile(path.Join(viper.GetString("configpath"), i.EnvConfiguration.Name+".yaml"))
}

func (c addRepositoryCommand) IsResponsible(commandName string) bool {
	return commandName == "addrepo"
}

func init() {
	commands = append(commands, addRepositoryCommand{})
}
