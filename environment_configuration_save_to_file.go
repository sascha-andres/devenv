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
	yaml "gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
)

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
