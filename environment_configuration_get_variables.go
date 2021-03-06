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
