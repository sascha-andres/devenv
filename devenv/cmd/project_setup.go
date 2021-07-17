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

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"path"
)

func projectSetup(projectName string) error {
	if "" == projectName || !devenv.ProjectIsCreated(projectName) {
		return errors.New(fmt.Sprintf("Project '%s' does not yet exist", projectName))
	}
	if ok, err := helper.Exists(path.Join(viper.GetString("configpath"), projectName+".yaml")); ok && err == nil {
		if err := ev.LoadFromFile(path.Join(viper.GetString("configpath"), projectName+".yaml")); err != nil {
			return errors.New(fmt.Sprintf("Error reading env config: '%s'", err.Error()))
		}
	}
	return nil
}
