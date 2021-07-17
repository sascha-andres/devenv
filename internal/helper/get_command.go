// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
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

// GetCommand creates a command structure
func GetCommand(commandName string, env os_helper.Environ, path string, arguments ...string) (*exec.Cmd, error) {
	command := exec.Command(commandName, arguments...)
	command.Dir = path
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return command, nil
}
