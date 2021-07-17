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
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

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
