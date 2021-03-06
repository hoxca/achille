/*
Copyright © 2020 Hugues Obolonsky <hugh@atosc.org>

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

// Package cmd for 'roof isOpen' Command
package cmd

import (
	"achille/roof"
	"fmt"

	"github.com/spf13/cobra"
)

// isOpenCmd represents the isOpen command
var isOpenCmd = &cobra.Command{
	Use:   "isOpen",
	Short: "achille command to confirm that talon roof is open",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		talonRoof := new(roof.Status)
		talonRoof = talonRoof.Set(splitedlog[3], splitedlog[2])

		if talonRoof.IsOpen() {
			fmt.Printf("OK, ")
			talonRoof.PrintStatus()
		} else {
			fmt.Printf("Nope!")
		}

	},
}

func init() {
	roofCmd.AddCommand(isOpenCmd)
}
