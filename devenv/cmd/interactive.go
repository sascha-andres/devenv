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

package cmd

import (
	"github.com/chzyer/readline"
	"github.com/sascha-andres/devenv"
)

var ev devenv.EnvironmentConfiguration

var completer = readline.NewPrefixCompleter(
	readline.PcItem("repo",
		readline.PcItemDynamic(listRepositories(),
			readline.PcItem("branch"),
			readline.PcItem("fetch"),
			readline.PcItem("commit"),
			readline.PcItem("log"),
			readline.PcItem("shell"),
			readline.PcItem("pull"),
			readline.PcItem("push"),
			readline.PcItem("status"),
			readline.PcItem("pin"),
			readline.PcItem("unpin"),
			readline.PcItem("enable"),
			readline.PcItem("disable"),
		),
	),
	readline.PcItem("addrepo"),
	readline.PcItem("addscript"),
	readline.PcItem("editscript"),
	readline.PcItem("delscript"),
	readline.PcItem("fetch"),
	readline.PcItem("branch"),
	readline.PcItem("commit"),
	readline.PcItem("delrepo"),
	readline.PcItem("log"),
	readline.PcItem("pull"),
	readline.PcItem("push"),
	readline.PcItem("status"),
	readline.PcItem("quit"),
	readline.PcItem("scan"),
	readline.PcItem("shell"),
)
