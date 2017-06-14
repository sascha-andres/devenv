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
	"html/template"
	"os/exec"
)

var (
	shPath   string
	shExists bool
)

// EnvironmentConfiguration contains information about the project
type (
	EnvironmentConfiguration struct {
		Name           string                    `yaml:"name"`
		Repositories   []RepositoryConfiguration `yaml:"repositories"`
		Environment    map[string]string         `yaml:"env"`
		Shell          string                    `yaml:"shell"`
		ShellArguments []string                  `yaml:"shell-arguments"`
		Commands       []string                  `yaml:"commands"`
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

func init() {
	var err error
	shPath, err = exec.LookPath("sh")
	if err != nil {
		return
	}
	shExists = true
}
