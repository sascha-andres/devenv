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
	"fmt"
	"log"
	"os"
	"os/exec"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sascha-andres/devenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	cfg        devenv.Configuration
	configPath string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "devenv",
	Short: "Manage your code in project environments",
	Long: `Manage your multi repository projects with ease.
Commit all repositories at once`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runInterpreter(args)
		if err != nil {
			log.Fatalf("Error: '%s'", err.Error)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("devenv version v1.2.0")
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.devenv.yaml)")
	RootCmd.PersistentFlags().StringP("basepath", "b", "$HOME/devenv/src", "Base path for projects")
	RootCmd.PersistentFlags().StringP("configpath", "c", "$HOME/devenv/environments", "Config path for environments")
	RootCmd.PersistentFlags().StringP("logconfiguration", "", "--oneline --graph --decorate --all", "Additional parameters for log calls")

	viper.BindPFlag("basepath", RootCmd.PersistentFlags().Lookup("basepath"))
	viper.BindPFlag("configpath", RootCmd.PersistentFlags().Lookup("configpath"))
	viper.BindPFlag("logconfiguration", RootCmd.PersistentFlags().Lookup("logconfiguration"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(home)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".devenv")
	}

	viper.SetEnvPrefix("DEVENV")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	path, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("No git installation found")
	}
	log.Printf("Using git at '%s'\n", path)
}
