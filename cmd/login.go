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
	"fssh/util"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login ssh",
	Long: ``,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {

			var currentNode Node
			for _, s := range ssh.Node {
				if s.Id == args[0] || strings.HasPrefix(s.Id, args[0]) || s.Name == args[0]{
					currentNode = s
					break
				}

				compile, err := regexp.Compile(`(\d+\.){3}\d+`)
				if err == nil{
					if compile.MatchString(args[0]){
						currentNode = s
						break
					}
				}
			}

			if currentNode.Host != ""{
				if currentNode.Name != ""{
					util.RunCommand("ssh", fmt.Sprintf("%s@%s", currentNode.User, currentNode.Host))
				}else{
					util.RunCommand("ssh", currentNode.Host)
				}
			}
		} else {
			fmt.Println("invalid argument")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
