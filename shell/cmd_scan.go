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

package shell

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/sascha-andres/devenv/helper"
	"github.com/spf13/viper"
)

type scanCommand struct{}

var (
	interpreter *Interpreter
)

func (c scanCommand) Execute(i *Interpreter, repository string, args []string) error {
	log.Println("Scanning for new repositories")
	interpreter = i
	err := filepath.Walk(i.ExecuteScriptDirectory, walker)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func walker(foundPath string, info os.FileInfo, err error) error {
	if info.IsDir() {
		localPath := foundPath
		if strings.HasPrefix(foundPath, interpreter.ExecuteScriptDirectory) {
			localPath = strings.Replace(foundPath, interpreter.ExecuteScriptDirectory+"/", "", 1)
		}
		if strings.HasSuffix(localPath, ".git") {
			return handleGitDirectory(localPath)
		}
	}
	return nil
}

func handleGitDirectory(localPath string) error {
	localPath = strings.Replace(localPath, ".git", "", 1)
	if !isKnownRepository(localPath) {
		return handleUnknownRepository(localPath)
	}
	return filepath.SkipDir
}

func handleUnknownRepository(localPath string) error {
	log.Printf("Unknown repository found at %s\n", localPath)
	log.Print("  Type y to add ")
	add := getAnswer()
	if strings.ToLower(add) == "y" {
		remote, err := getRemote(localPath)
		if remote == "" || err != nil {
			log.Println("Cannot use empty remote, probably local only?")
			return filepath.SkipDir
		}
		return addConfiguration(localPath, "", localPath, remote, false)
	}
	return filepath.SkipDir
}

func addConfiguration(localPath, pinned, name, url string, disabled bool) error {
	cfg := devenv.RepositoryConfiguration{
		Path:     localPath,
		Pinned:   pinned,
		Disabled: disabled,
		Name:     name,
		URL:      url,
	}
	log.Println(cfg)
	interpreter.EnvConfiguration.Repositories = append(interpreter.EnvConfiguration.Repositories, cfg)
	return interpreter.EnvConfiguration.SaveToFile(path.Join(viper.GetString("configpath"), interpreter.EnvConfiguration.Name+".yaml"))
}

func getRemote(localPath string) (string, error) {
	repoPath := path.Join(interpreter.ExecuteScriptDirectory, localPath)
	var arguments []string
	arguments = append(arguments, "remote", "get-url", "origin")
	vars, err := interpreter.EnvConfiguration.GetReplacedEnvironment()
	if err != nil {
		return "", filepath.SkipDir
	}
	var remote string
	if remote, err = helper.GitOutput(vars, repoPath, arguments...); err != nil {
		return "", filepath.SkipDir
	}
	return remote, nil
}

func isKnownRepository(relativePath string) bool {
	for _, repository := range interpreter.EnvConfiguration.Repositories {
		if relativePathIsMatch(relativePath, repository.Path) {
			return true
		}
	}
	return false
}

func relativePathIsMatch(relativePath, compareTo string) bool {
	localRelativePath := strings.Replace(relativePath+"/", "//", "/", -1)
	localCompareTo := strings.Replace(compareTo+"/", "//", "/", -1)
	return localRelativePath == localCompareTo
}

func (c scanCommand) IsResponsible(commandName string) bool {
	return commandName == "scan"
}

func init() {
	commands = append(commands, scanCommand{})
}
