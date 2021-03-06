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
