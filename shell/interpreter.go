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
	"log"

	"github.com/mgutz/str"
)

type (
	// Interpreter contains data for where and what to do
	Interpreter struct {
		ExecuteScriptDirectory string
	}
)

// NewInterpreter returns a new interpreter
func NewInterpreter(path string) *Interpreter {
	return &Interpreter{ExecuteScriptDirectory: path}
}

func (i *Interpreter) Execute(commandline string) error {
	tokenized := str.ToArgv(commandline)
	log.Printf("%#v\n", tokenized)
	switch tokenized[0] {
	case "repo":
		log.Println("Repository specific function")
	case "branch":
		log.Println("Create branch in all repositories")
	case "pull":
		log.Println("Get latest code for all repositories")
	case "push":
		log.Println("Put latest code")
	}
	return nil
}
