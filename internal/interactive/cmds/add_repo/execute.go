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

package add_repo

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"log"
	"path"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repository string, args []string) error {
	log.Println("Add a repo")
	log.Print("Please provide a link: ")
	link := helper.GetAnswer()
	log.Print("Please provide a name: ")
	name := helper.GetAnswer()
	log.Print("Please provide a relative path: ")
	rpath := helper.GetAnswer()
	repositoryPath := path.Join(executeScriptDirectory, rpath)
	vars, err := e.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	if _, err := helper.Git(vars, executeScriptDirectory, "clone", link, repositoryPath); err != nil {
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
	e.Repositories = append(e.Repositories, rc)
	return e.SaveToFile(path.Join(viper.GetString("configpath"), e.Name+".yaml"))
}
