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

package repo_log

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"path"
	"strings"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repository string, args []string) error {
	_, repo := e.GetRepository(repository)
	repositoryPath := path.Join(executeScriptDirectory, repo.Path)
	vars, err := e.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	var params = []string{"log"}
	params = append(params, strings.Split(viper.GetString("logconfiguration"), " ")...)
	params = append(params, args...)
	_, err = helper.Git(vars, repositoryPath, params...)
	return err
}
