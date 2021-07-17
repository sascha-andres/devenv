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

package helper

import (
	"github.com/sascha-andres/devenv/internal/os_helper"
	"os"
	"os/exec"
)

var (
	gitExecutable string
)

// Git calls the system git in the project directory with specified arguments
func Git(ev map[string]string, projectPath string, args ...string) (int, error) {
	command := exec.Command(gitExecutable, args...)
	env := BuildEnvironment(ev)
	command.Dir = projectPath
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return os_helper.StartAndWait(command)
}
