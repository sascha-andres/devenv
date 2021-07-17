// Licensed under the Apache License, Version 2.0 (the "License");
// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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

package helper

// HasBranch checks if there is a branch locally or remotely
func HasBranch(ev map[string]string, projectPath, branch string) (bool, error) {
	ok, err := HasRemoteBranch(ev, projectPath, branch)
	if err != nil {
		return false, err
	}
	ok2, err := HasLocalBranch(ev, projectPath, branch)
	return ok && ok2, err
}
