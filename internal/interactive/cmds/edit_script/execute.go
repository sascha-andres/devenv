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

package edit_script

import (
	"github.com/pkg/errors"
	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/internal/helper"
	"github.com/sascha-andres/devenv/internal/os_helper"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
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
		log.Printf("error: script does not exist")
		return errors.New("no such script")
	}

	environmentVariables := os_helper.GetEnvironmentVariables()
	value, found := environmentVariables["EDITOR"]
	if !found {
		log.Println("no EDITOR variable found")
		return errors.New("no EDITOR variable found")
	}
	command := exec.Command(value, []string{path.Join(viper.GetString("basepath"), e.Name, scriptName)}...)
	env := helper.BuildEnvironment(environmentVariables)
	command.Dir = path.Join(viper.GetString("basepath"), e.Name)
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	_, err := os_helper.StartAndWait(command)
	return err
}
