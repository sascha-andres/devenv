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
	"fmt"
	"github.com/sascha-andres/devenv/internal/os_helper"
	"github.com/spf13/viper"
	"log"
	"path"
	"time"
)

//PushConfiguration puts  the latest commit to remote repository
func PushConfiguration() error {
	configPath := viper.GetString("configpath")
	if ok, _ := os_helper.Exists(path.Join(configPath, ".git")); !ok {
		return nil
	}

	cmd, err := os_helper.GetCommand("git", nil, configPath, "add", "--all", ":/")
	if err != nil {
		return err
	}
	_, err = os_helper.StartAndWait(cmd)
	if err != nil {
		log.Println("could not add all changes to repository")
		return err
	}

	commitMessage := fmt.Sprintf("%s: changes to devenv configuration", time.Now().Format(time.RFC3339))

	cmd, err = os_helper.GetCommand("git", nil, configPath, "commit", "-m", fmt.Sprintf("\"%s\"", commitMessage))
	if err != nil {
		log.Println("could not add commit to git ")
		return err
	}
	_, err = os_helper.StartAndWait(cmd)
	if err != nil {
		log.Println("could not add all changes to repository")
		return err
	}

	cmd, err = os_helper.GetCommand("git", nil, configPath, "push")
	if err != nil {
		log.Println("could not push to remote server")
		return err
	}

	_, err = os_helper.StartAndWait(cmd)
	return err
}
