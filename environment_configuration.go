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
	"bytes"
	"fmt"
	"github.com/sascha-andres/devenv/internal/os_helper"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

//var (
//  shPath   string
//  shExists bool
//)

type (
	// EnvironmentConfiguration contains information about the project
	EnvironmentConfiguration struct {
		Name                 string                                  `yaml:"name"`
		Repositories         []RepositoryConfiguration               `yaml:"repositories"`
		ProcessConfiguration EnvironmentExternalProcessConfiguration `yaml:"processes"`
	}

	// EnvironmentExternalProcessConfiguration contains configuration in use with external processes
	EnvironmentExternalProcessConfiguration struct {
		Environment    map[string]string `yaml:"env"`
		Shell          string            `yaml:"shell"`
		ShellArguments []string          `yaml:"shell-arguments"`
		Commands       []string          `yaml:"commands"`
	}
)

// applyVariables uses GO's templating to apply the variables
func (ev *EnvironmentConfiguration) applyVariables(input string) (string, error) {
	templ, err := template.New("").Parse(input)
	if err != nil {
		return "", err
	}
	b := bytes.NewBuffer(nil)
	vars, err := ev.GetVariables()
	if err != nil {
		return "", err
	}
	if err = templ.Execute(b, vars); err != nil {
		return "", err
	}
	return b.String(), nil
}

// GetEnvironment returns the templated env vars
func (ev *EnvironmentConfiguration) GetEnvironment() ([]string, error) {
	env := os_helper.Environ(os.Environ())
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

// GetReplacedEnvironment provides a way to get the environment variables with replaced values
func (ev *EnvironmentConfiguration) GetReplacedEnvironment() (map[string]string, error) {
	var localvars map[string]string
	for key, val := range ev.ProcessConfiguration.Environment {
		result, err := ev.applyVariables(val)
		if err != nil {
			return nil, err
		}
		localvars[key] = result
	}

	return localvars, nil
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

// GetVariables returns variable map for environment
func (ev *EnvironmentConfiguration) GetVariables() (map[string]string, error) {
	localVariables := make(map[string]string)
	for key, value := range os_helper.GetEnvironmentVariables() {
		localVariables[key] = value
	}
	for key, value := range ev.ProcessConfiguration.Environment {
		localVariables[key] = value
	}
	localVariables["ENV_DIRECTORY"] = path.Join(viper.GetString("basepath"), ev.Name)
	return localVariables, nil
}

// LoadFromFile takes a YAML file and unmarshals its data
func (ev *EnvironmentConfiguration) LoadFromFile(path string) error {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening project config: %#v\n", err)
	}
	file, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error loading project config: %#v\n", err)
	}
	if err := yaml.Unmarshal(file, ev); err != nil {
		log.Fatalf("Error reading project config: %#v\n", err)
	}
	return nil
}

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

// ProjectIsCreated checks whether project is checked out
func ProjectIsCreated(projectName string) bool {
	if ok, err := os_helper.Exists(path.Join(viper.GetString("basepath"), projectName)); ok && err == nil {
		return true
	}
	return false
}

// RepositoryExists returns true if a repository is configured in environment
func (ev *EnvironmentConfiguration) RepositoryExists(repoName string) bool {
	_, repo := ev.GetRepository(repoName)
	return repo != nil
}

// SaveToFile takes the config and saves to disk
func (ev *EnvironmentConfiguration) SaveToFile(path string) error {
	data, err := yaml.Marshal(ev)
	if err != nil {
		log.Fatalf("Error marshalling project config: %#v\n", err)
	}
	err = ioutil.WriteFile(path, data, 0600)
	if err != nil {
		log.Fatalf("Error saving project config: %#v\n", err)
	}
	return nil
}

// StartShell executes configured shell or default shell (sh)
func (ev *EnvironmentConfiguration) StartShell() error {
	return ev.StartShellForSubdirectory("")
}

// StartShellForSubdirectory executes configured shell or default shell (sh) within a specified subdirectory
func (ev *EnvironmentConfiguration) StartShellForSubdirectory(subdirectory string) error {
	err := ev.prepareShell()
	if err != nil {
		return err
	}

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
	command, err := os_helper.GetCommand(commandName, env, path.Join(viper.GetString("basepath"), ev.Name, subdirectory), arguments...)
	if err != nil {
		return err
	}
	_, err = os_helper.StartAndWait(command)
	if err != nil {
		return err
	}
	return nil
}
