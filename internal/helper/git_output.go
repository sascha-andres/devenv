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
	"bytes"
	"os/exec"
	"strings"
)

// GitOutput calls the system git in the project directory with specified arguments and returns the output
func GitOutput(ev map[string]string, projectPath string, args ...string) (string, error) {
	command := exec.Command(gitExecutable, args...)
	env := BuildEnvironment(ev)
	command.Dir = projectPath
	command.Env = env
	stdout, err := command.StdoutPipe()
	err = command.Start()
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	output := buf.String()
	return strings.TrimSpace(output), err
}
