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

package interactive

import (
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
)

type addScriptCommand struct{}

func (c addScriptCommand) Execute(i *Interpreter, repositoryName string, args []string) error {
	var filePath = strings.Join(args, "")
	if ok, err := helper.Exists(filePath); !ok {
		log.Printf("[%s] does not exist", filePath)
		return err
	}
	var scriptDirectory = path.Join(viper.GetString("configpath"), i.EnvConfiguration.Name)
	if ok, _ := helper.Exists(scriptDirectory); !ok {
		err := os.MkdirAll(scriptDirectory, 0700)
		if err != nil {
			log.Printf("Could not create %s", scriptDirectory)
			return err
		}
	}
	// TODO copy file to script directory and make executable (0700)
	// TODO on linux/osx symlink script in env dir to script directory
	// TODO on Windows copy script tp env dir

	return nil
}

func (c addScriptCommand) IsResponsible(commandName string) bool {
	return commandName == "addscript"
}

func init() {
	commands = append(commands, addScriptCommand{})
}
