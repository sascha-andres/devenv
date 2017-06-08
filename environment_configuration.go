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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"

	yaml "gopkg.in/yaml.v1"
)

var (
	shPath   string
	shExists bool
)

// EnvironmentConfiguration contains information aout the project
type (
	EnvironmentConfiguration struct {
		Name         string                    `yaml:"name"`
		BranchPrefix string                    `yaml:"branch-prefix"`
		Repositories []RepositoryConfiguration `yaml:"repositories"`
		Environment  map[string]string         `yaml:"env"`
		Shell        string                    `yaml:"shell"`
		Commands     []string                  `yaml:"commands"`
	}
)

// LoadFromFile takes a YAML file and unmarhals its data
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
func (ev *EnvironmentConfiguration) prepareShell() error {
	for _, cmd := range ev.Commands {
		var command *exec.Cmd
		command = exec.Command("bash", "-l", "-c", cmd)
		env := helper.Environ(os.Environ())
		for key := range ev.Environment {
			env.Unset(key)
		}
		for key, value := range ev.Environment {
			log.Printf("Setting '%s' to '%s'", key, value)
			env = append(env, fmt.Sprintf("%s=%s", key, value))
		}
		command.Dir = path.Join(viper.GetString("basepath"), ev.Name)
		command.Env = env
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		if err := command.Start(); err != nil {
			log.Printf("Error executing '%s': '%s'", cmd, err.Error())
		}
		if err := command.Wait(); err != nil {
			log.Printf("Error executing '%s': '%s'", cmd, err.Error())
		}
	}
	return nil
}

// StartShell executes configured shell or default shell (sh)
func (ev *EnvironmentConfiguration) StartShell() error {
	ev.prepareShell()
	var command *exec.Cmd
	if ev.Shell != "" {
		command = exec.Command(ev.Shell)
	} else {
		command = exec.Command("bash", "-l")
	}
	env := helper.Environ(os.Environ())
	for key := range ev.Environment {
		env.Unset(key)
	}
	for key, value := range ev.Environment {
		log.Printf("Setting '%s' to '%s'", key, value)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	command.Dir = path.Join(viper.GetString("basepath"), ev.Name)
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		return fmt.Errorf("Error running bash: %#v", err)
	}
	if err := command.Wait(); err != nil {
		return fmt.Errorf("Error waiting for bash: %#v", err)
	}
	return nil
}

// RepositoryExists returns true if a repository is configured in environment
func (ev *EnvironmentConfiguration) RepositoryExists(repoName string) bool {
	for _, repo := range ev.Repositories {
		if repo.Name == repoName {
			return true
		}
	}
	return false
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

func init() {
	var err error
	shPath, err = exec.LookPath("sh")
	if err != nil {
		return
	}
	shExists = true
}
