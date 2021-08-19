/*
Copyright © 2021 ConsenSys Software Inc.

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
)

// groth16Cmd represents the groth16 command
var groth16Cmd = &cobra.Command{
	Use:   "groth16",
	Short: "runs benchmarks and profiles using Groth16 proof system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("groth16 called")
	},
}

func init() {
	rootCmd.AddCommand(groth16Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groth16Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groth16Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
