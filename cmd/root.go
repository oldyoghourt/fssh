/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type SSH struct {
	Node []Node
}

type Node struct {
	Id     string
	Name   string
	User   string
	Passwd string
	Host   string
}

var cfgFile string
var ssh SSH

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fssh",
	Short: "fast login ssh",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/fsshrc.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "fsshrc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName("fsshrc")
	}

	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		err := viper.Unmarshal(&ssh)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

// search the node, return the index and Node
func SelectNode(args []string, f func([]string) (int, Node)) (int, Node) {
	var choiceNode Node
	i := -1
	if len(args) == 0 {
		return i, choiceNode
	}
	if nil != f {
		return f(args)
	}

	for index, s := range ssh.Node {
		if s.Id == args[0] || strings.HasPrefix(s.Id, args[0]) || s.Name == args[0] {
			choiceNode = s
			i = index
			break
		}

		compile, err := regexp.Compile(`(\d+\.){3}\d+`)
		if err == nil {
			if compile.MatchString(args[0]) && s.Host == args[0] {
				choiceNode = s
				i = index
				break
			}
		}
	}
	return i, choiceNode
}
