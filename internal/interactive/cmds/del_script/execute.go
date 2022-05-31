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

package del_script

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func (c Command) Execute(e *devenv.EnvironmentConfiguration, executeScriptDirectory, repositoryName string, args []string) error {
	_ = executeScriptDirectory

	if !(runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		log.Printf("editscript is not supported on %s", runtime.GOOS)
		return nil
	}
	scriptName := strings.Join(args, "")
	if ok, _ := helper.Exists(path.Join(viper.GetString("basepath"), e.Name, scriptName)); !ok {
		log.Printf("error: script does not exist in environment folder")
		return errors.New("no such script in environment folder")
	}

	var scriptDirectory = path.Join(viper.GetString("configpath"), e.Name)
	var scriptInScriptDirectory = path.Join(scriptDirectory, scriptName)
	if ok, _ := helper.Exists(scriptInScriptDirectory); !ok {
		log.Printf("error: script does not exist in config directory")
		return errors.New(fmt.Sprintf("no such script in configuration folder %s", scriptDirectory))
	}

	err := os.Remove(scriptInScriptDirectory)
	if err != nil {
		return err
	}

	return os.RemoveAll(path.Join(viper.GetString("basepath"), e.Name, scriptName))
}
