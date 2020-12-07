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
	"github.com/hashicorp/go-uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add host info",
	Long: `
	HOST INFO:
	-	Name		node name
	-	User		user name
	-	Host		host address
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var n Node
		n.Id, _ = uuid.GenerateUUID()
		cmd.Print("Please input Name: ")
		_, _ = fmt.Scanln(&n.Name)
		cmd.Print("Please input User: ")
		_, _ = fmt.Scanln(&n.User)
		//fmt.Print("Please input Passwd: ")
		//_, _ = fmt.Scanln(&n.Passwd)
		//aes, err := util.EncryptAES([]byte(n.Passwd), []byte(secret))
		//if err != nil{
		//	fmt.Println(err.Error())
		//}
		//n.Passwd = string(aes)
		cmd.Print("Please input Host: ")
		_, _ = fmt.Scanln(&n.Host)
		ssh.Node = append(ssh.Node, n)
		viper.Set("node", ssh.Node)
		err := viper.SafeWriteConfig()
		if err != nil {
			if strings.Contains(err.Error(), "Already Exists") {
				_ = viper.WriteConfig()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
