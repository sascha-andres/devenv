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

package repo_pin

import (
	"github.com/sascha-andres/devenv"
	helper2 "github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"log"
	"path"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repositoryName string, args []string) error {
	index, repository := e.GetRepository(repositoryName)
	log.Printf("Pinning %s", repository.Name)

	repositoryPath := path.Join(executeScriptDirectory, repository.Path)
	vars, err := e.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	var params = []string{"rev-parse"}
	params = append(params, "--verify", "HEAD")
	output, err := helper2.GitOutput(vars, repositoryPath, params...)
	if err != nil {
		return err
	}
	repository.Pinned = output
	e.Repositories = append(e.Repositories[:index], e.Repositories[index+1:]...)
	e.Repositories = append(e.Repositories, *repository)
	return e.SaveToFile(path.Join(viper.GetString("configpath"), e.Name+".yaml"))
	return nil
}
