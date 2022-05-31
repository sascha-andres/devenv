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

package interactive

func tryExecuteCommand(val Commander, i *Interpreter, specific bool, arguments []string) (bool, error) {
	commandIndex := 0
	if specific {
		commandIndex = 1
	}
	if len(arguments)+1 > commandIndex {
		return false, nil
	}
	if val.IsResponsible(arguments[commandIndex]) {
		if specific {
			return true, val.Execute(&i.EnvConfiguration, i.ExecuteScriptDirectory, arguments[0], arguments[2:])
		}
		return true, val.Execute(&i.EnvConfiguration, i.ExecuteScriptDirectory, "%", arguments[1:])
	}
	return false, nil
}
