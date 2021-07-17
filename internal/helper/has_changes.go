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

package helper

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
)

// HasChanges checks whether a repo is clean or has changes ( modifications or additions )
func HasChanges(ev map[string]string, projectPath string) (bool, error) {
	_, err := os.Stat(projectPath)
	if os.IsNotExist(err) {
		return false, nil
	}
	// git status --porcelain
	command := exec.Command(gitExecutable, "status", "--porcelain")
	env := BuildEnvironment(ev)
	command.Dir = projectPath
	command.Env = env
	out, err := command.Output()
	if err != nil {
		return true, err
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for scanner.Scan() {
		return true, nil
	}
	return false, nil
}
