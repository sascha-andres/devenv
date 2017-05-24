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

package shell

import (
	"fmt"
	"log"
)

// ExecuteRepo checks command and hands over to specific functions
func (i *Interpreter) ExecuteRepo(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No repo and action provided")
	}
	if len(args) == 1 {
		return fmt.Errorf("No action provided")
	}
	if !i.EnvConfiguration.RepositoryExists(args[0]) {
		return fmt.Errorf("Repository unknown")
	}
	switch args[1] {
	case "lg":
		fallthrough
	case "log":
		return i.executeRepoLog(args[0])
	case "st":
		fallthrough
	case "status":
		return i.executeRepoStatus(args[0])
	case ">":
		fallthrough
	case "push":
		return i.executeRepoPush(args[0], args[2:])
	case "<":
		fallthrough
	case "pull":
		return i.executeRepoPush(args[0], args[2:])
	default:
		log.Printf("%v\n", args)
	}
	return nil
}
