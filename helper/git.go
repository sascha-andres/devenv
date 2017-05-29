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

package helper

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	gitExecutable string
)

// Git calls the system git in the project diorectory with specified arguments
func Git(ev map[string]string, projectPath string, args ...string) error {
	command := exec.Command("git", args...)
	env := Environ(os.Environ())
	for key := range ev {
		env.Unset(key)
	}
	for key, value := range ev {
		log.Printf("Setting '%s' to '%s'", key, value)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	command.Dir = projectPath
	command.Env = env
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		return fmt.Errorf("Error running bash: %#v", err)
	}
	if err := command.Wait(); err != nil {
		return fmt.Errorf("Error waiting for bash: %#v", err)
	}
	return nil
}

func init() {
	var err error
	gitExecutable, err = exec.LookPath("git")
	if err != nil {
		log.Fatalf("Could not locate git: '%#v'", err)
	}
}

// HasChanges checks whether a repo is clean or has changes ( modifications or additions )
func HasChanges(ev map[string]string, projectPath string) (bool, error) {
	// git status --porcelain
	command := exec.Command("git", "status", "--porcelain")
	env := Environ(os.Environ())
	for key := range ev {
		env.Unset(key)
	}
	for key, value := range ev {
		log.Printf("Setting '%s' to '%s'", key, value)
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	command.Dir = projectPath
	command.Env = env
	out, err := command.Output()
	if err != nil {
		return true, err
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for scanner.Scan() {
		return true, nil
	}
	return false, nil
}

// HasBranch checks if there is a branch locally or remotely
func HasBranch(ev map[string]string, projectPath, branch string) (bool, error) {
	ok, err := HasRemoteBranch(ev, projectPath, branch)
	if err != nil {
		return false, err
	}
	ok2, err := HasLocalBranch(ev, projectPath, branch)
	return ok && ok2, err
}

// HasRemoteBranch checks if there is a branch remotely
func HasRemoteBranch(ev map[string]string, projectPath, branch string) (bool, error) {
	// git ls-remote --exit-code . origin/branch-name -> 0 == remote existiert
	return false, nil
}

// HasLocalBranch checks if there is a branch locally
func HasLocalBranch(ev map[string]string, projectPath, branch string) (bool, error) {
	// git rev-parse --verify <branch> --> 0 == local branch exists
	return false, nil
}
