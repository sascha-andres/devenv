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
	"log"
	"os"
	"path"
	"strings"

	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean project environment",
	Long: `Removes all code from your harddisk. Scans all repositories for uncommitted
changes before deleting the complete directory tree.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := strings.Join(args, " ")
		log.Printf("Called to clean '%s'\n", projectName)
		log.Println("Currently directory is going to be removed withour any checks!")
		if devenv.ProjectIsCreated(projectName) {
			os.RemoveAll(path.Join(viper.GetString("basepath"), projectName))
		} else {
			log.Println("Project does not exist")
		}
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
