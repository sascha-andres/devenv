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
	"path"

	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
)

type addRepoCommand struct{}

func (c addRepoCommand) Execute(i *Interpreter, repository string, args []string) error {
	fmt.Println("Add a repo")
	fmt.Print("Please provide a link: ")
	link := getAnswer()
	fmt.Print("Please provide a name: ")
	name := getAnswer()
	fmt.Print("Please provide a relative path: ")
	rpath := getAnswer()
	repoPath := path.Join(i.ExecuteScriptDirectory, rpath)
	if _, err := helper.Git(i.EnvConfiguration.Environment, i.ExecuteScriptDirectory, "clone", link, repoPath); err != nil {
		log.Fatalf("Error executing git: '%#v'", err)
	}
	fmt.Printf("Repository '%s' saved as '%s' to relative path '%s'\n", link, name, rpath)
	rc := devenv.RepositoryConfiguration{Name: name, Path: rpath, URL: link}
	i.EnvConfiguration.Repositories = append(i.EnvConfiguration.Repositories, rc)
	return i.EnvConfiguration.SaveToFile(path.Join(viper.GetString("configpath"), i.EnvConfiguration.Name+".yaml"))
}

func (c addRepoCommand) IsResponsible(commandName string) bool {
	return commandName == "addrepo"
}

func init() {
	commands = append(commands, addRepoCommand{})
}
