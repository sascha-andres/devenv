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

package interactive

import (
	"github.com/sascha-andres/devenv/helper"
)

func execHelper(i *Interpreter, repoPath, command string, args []string) error {
	var arguments []string
	arguments = append(arguments, command)
	arguments = append(arguments, args...)
	vars, err := i.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return err
	}
	if _, err = helper.Git(vars, repoPath, arguments...); err != nil {
		return err
	}
	return nil
}
