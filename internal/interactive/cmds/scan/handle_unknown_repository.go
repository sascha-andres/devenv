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

package scan

import (
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"log"
	"path/filepath"
	"strings"
)

func handleUnknownRepository(localPath string) error {
	log.Printf("Unknown repository found at %s\n", localPath)
	log.Print("  Type y to add ")
	add := helper.GetAnswer()
	if strings.ToLower(add) == "y" {
		remote, err := getRemote(localPath)
		if remote == "" || err != nil {
			log.Println("Cannot use empty remote, probably local only?")
			return filepath.SkipDir
		}
		return addConfiguration(devenv.RepositoryConfiguration{
			Path:     localPath,
			Pinned:   "",
			Disabled: false,
			Name:     remote,
			URL:      remote,
		})
	}
	return filepath.SkipDir
}
