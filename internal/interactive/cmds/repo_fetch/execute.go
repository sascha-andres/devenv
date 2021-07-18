// Licensed under the Apache License, Version 2.0 (the "License");
// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
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

package repo_fetch

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"path"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repositoryName string, args []string) error {
	_, repository := e.GetRepository(repositoryName)
	if repository.Disabled || repository.Pinned != "" {
		return nil
	}

	repositoryPath := path.Join(executeScriptDirectory, repository.Path)
	return helper.ExecHelper(e, repositoryPath, "fetch", args)
}
