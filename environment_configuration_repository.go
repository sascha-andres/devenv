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

package devenv

import (
	"path"

	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
)

// RepositoryExists returns true if a repository is configured in environment
func (ev *EnvironmentConfiguration) RepositoryExists(repoName string) bool {
	_, repo := ev.GetRepository(repoName)
	return repo != nil
}

// GetRepository returns a repository with given name
func (ev *EnvironmentConfiguration) GetRepository(repoName string) (int, *RepositoryConfiguration) {
	for index, repo := range ev.Repositories {
		if repo.Name == repoName {
			return index, &repo
		}
	}
	return 0, nil
}

// ProjectIsCreated checks whether project is checked out
func ProjectIsCreated(projectName string) bool {
	if ok, err := helper.Exists(path.Join(viper.GetString("basepath"), projectName)); ok && err == nil {
		return true
	}
	return false
}
