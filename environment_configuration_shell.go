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
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
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
		command, err := helper.GetCommand(commandName, env, path.Join(viper.GetString("basepath"), ev.Name), arguments...)
		if err != nil {
			return err
		}
		_, err = helper.StartAndWait(command)
		if err != nil {
			return err
		}
	}
	return nil
}

// StartShell executes configured shell or default shell (sh)
func (ev *EnvironmentConfiguration) StartShell() error {
	ev.prepareShell()
	commandName, arguments, err := ev.GetShell()
	if err != nil {
		return err
	}
	for _, val := range ev.ProcessConfiguration.ShellArguments {
		result, err := ev.applyVariables(val)
		if err != nil {
			return err
		}
		arguments = append(arguments, result)
	}
	env, err := ev.GetEnvironment()
	if err != nil {
		return err
	}
	command, err := helper.GetCommand(commandName, env, path.Join(viper.GetString("basepath"), ev.Name), arguments...)
	if err != nil {
		return err
	}
	_, err = helper.StartAndWait(command)
	if err != nil {
		return err
	}
	return nil
}

// GetEnvironment returns the templated env vars
func (ev *EnvironmentConfiguration) GetEnvironment() ([]string, error) {
	env := helper.Environ(os.Environ())
	for key := range ev.ProcessConfiguration.Environment {
		env.Unset(key)
	}
	for key, value := range ev.ProcessConfiguration.Environment {
		result, err := ev.applyVariables(value)
		if err != nil {
			return nil, err
		}
		log.Printf("Setting '%s' to '%s'", key, result)
		env = append(env, fmt.Sprintf("%s=%s", key, result))
	}
	return env, nil
}

// GetShell returns the shell executable with template applied
func (ev *EnvironmentConfiguration) GetShell() (string, []string, error) {
	if ev.ProcessConfiguration.Shell != "" {
		result, err := ev.applyVariables(ev.ProcessConfiguration.Shell)
		if err != nil {
			return "", nil, err
		}
		return result, nil, nil
	}
	return "bash", []string{"-l"}, nil
}
