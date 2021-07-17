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

package devenv

import (
	"github.com/sascha-andres/devenv/internal/os_helper"
	"github.com/spf13/viper"
	"path"
)

// prepareShell executes configured pre shell commands
func (ev *EnvironmentConfiguration) prepareShell() error {
	for _, cmd := range ev.ProcessConfiguration.Commands {
		commandName := "bash"
		result, err := ev.applyVariables(cmd)
		if err != nil {
			return err
		}
		arguments := []string{"-l", "-c", result}
		env, err := ev.GetEnvironment()
		if err != nil {
			return err
		}
		command, err := os_helper.GetCommand(commandName, env, path.Join(viper.GetString("basepath"), ev.Name), arguments...)
		if err != nil {
			return err
		}
		_, err = os_helper.StartAndWait(command)
		if err != nil {
			return err
		}
	}
	return nil
}
