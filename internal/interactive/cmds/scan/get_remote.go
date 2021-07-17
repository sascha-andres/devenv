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

package scan

import (
	helper2 "github.com/sascha-andres/devenv/internal/helper"
	"path"
	"path/filepath"
)

func getRemote(localPath string) (string, error) {
	repoPath := path.Join(packageExecuteScriptDirectory, localPath)
	var arguments []string
	arguments = append(arguments, "remote", "get-url", "origin")
	vars, err := envConfig.GetReplacedEnvironment()
	if err != nil {
		return "", filepath.SkipDir
	}
	var remote string
	if remote, err = helper2.GitOutput(vars, repoPath, arguments...); err != nil {
		return "", filepath.SkipDir
	}
	return remote, nil
}
