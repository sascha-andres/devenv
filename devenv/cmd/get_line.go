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

package cmd

import (
	"github.com/chzyer/readline"
	"io"
)

func getLine(l *readline.Instance) (string, bool) {
	line, err := l.Readline()
	if err == readline.ErrInterrupt {
		if len(line) == 0 {
			return "", true
		}
		return line, false
	} else if err == io.EOF {
		return "", true
	}
	return line, false
}
