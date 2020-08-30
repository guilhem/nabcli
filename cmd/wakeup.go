/*
Copyright © 2020 Guilhem Lettron <guilhem@barpilot.io>

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
	"net"
	"os"

	"github.com/guilhem/nabcli/pkg/client"
	"github.com/spf13/cobra"
)

// wakeupCmd represents the wakeup command
var wakeupCmd = &cobra.Command{
	Use:   "wakeup",
	Short: "wakeup your nabaztag",

	Run: func(cmd *cobra.Command, args []string) {
		conn, err := net.Dial("tcp", "localhost:10543")
		if err != nil {
			fmt.Printf("Failed to connect to nabd: %s\n", err)
			os.Exit(1)
		}
		defer conn.Close()

		c := client.New(conn)
		if err := c.Wakeup(); err != nil {
			fmt.Printf("Failed to wakeup: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(wakeupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wakeupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wakeupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
